package SherryDB

import (
   "strconv"
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "encoding/json"
   "reflect"
   "bytes"
   "fmt"
)

// 資料庫連線設定
type DBConnect struct {
   DBMSType	string	// 資料庫類別, ex:MySQL
   DBMS     string   // 資料庫管理系統名稱（皆小寫）
   DbServer string   // 資料庫主機
   DbPort   string   // 資料庫主機對外服務port
   DbName   string   // 資料庫名稱
   DbLogin  string   // 登入帳號
   DbPasswd string   // 登入密碼
}

type MySQL struct {
   Config DBConnect
   Name string
   Conn *sql.DB
}

func(m *MySQL) CheckAndReConnect() {
   if m.Conn.Ping() != nil {
      _ = m.Database(m.Config)
   }
}

// 取得單筆資料
func (m *MySQL) DoreSelOne(sql string, t interface{}, cond ...interface{}) (val interface{}, err error) {
    m.CheckAndReConnect()
    s := reflect.ValueOf(t).Elem()

    onerow := make([]interface{}, s.NumField())
    for i := 0; i < s.NumField(); i++ {
        onerow[i] = s.Field(i).Addr().Interface()
    }

    stmt, err := m.Conn.Prepare(sql)
    if err != nil {
        return nil, err
    }
    defer stmt.Close()
    row := stmt.QueryRow(cond...)
    err = row.Scan(onerow...)
    if err != nil {
        return nil, err
    }
    return t, nil
}

// 取得多筆資料
// 採用SQL直接執行，是較不安全的作法
func (m *MySQL) DoreFetchHash(sqlString string) (string, error) {
   m.CheckAndReConnect()
   tableData := make([]map[string]interface{}, 0)

   rows, err := m.Conn.Query(sqlString)
   if err != nil {
      return "", fmt.Errorf("Exec SQL error:%v", err)
   }
   defer rows.Close()
   columns, err := rows.Columns()
   if err != nil {
       return "", err
   }
   count := len(columns)
   values := make([]interface{}, count)
   valuePtrs := make([]interface{}, count)
   for rows.Next() {
       for i := 0; i < count; i++ {
           valuePtrs[i] = &values[i]
       }
       rows.Scan(valuePtrs...)
       entry := make(map[string]interface{})
       for i, col := range columns {
           var v interface{}
           val := values[i]
           b, ok := val.([]byte)
           if ok {
               v = string(b)
           } else {
               v = val
           }
           entry[col] = v
       }
       tableData = append(tableData, entry)
   }
   jsonData, err := json.Marshal(tableData)
   if err != nil {
       return "", err
   }
   return string(jsonData), nil 
}

func (m *MySQL) SelMultiple(sql string, t interface{}, cond ...interface{}) (*[]interface{}, error) {
   m.CheckAndReConnect()
   stmt, err := m.Conn.Prepare(sql)
   if err != nil {
      return nil, err
   }
   defer stmt.Close()
   rows, err := stmt.Query(cond...)
   if err != nil {
      return nil, err
   }
   defer rows.Close()

   vals := make([]interface{}, 0)
   s := reflect.ValueOf(t).Elem()
   row := make([]interface{}, s.NumField())
   for i := 0; i < s.NumField(); i++ {
      row[i] = s.Field(i).Addr().Interface()
   }
   for rows.Next() {
      err = rows.Scan(row...)
      if err != nil {
         return nil, err
      }
      vals = append(vals, s.Interface())
   }
   return &vals, nil
}

// 執行SQL指令並回傳新ID
func (m *MySQL) Insert(sqlstr string, cond ...string) (string, error) {
   m.CheckAndReConnect()
   stmt, err := m.Conn.Prepare(sqlstr)
   if err != nil {
      return "", fmt.Errorf("Prepare SQL error: %v", err)
   }
   t := make([]interface{}, len(cond))
   for i, val := range cond {
      t[i] = val
   }
   res, err := stmt.Exec(t)
   defer stmt.Close()
   if err != nil {
      return "", err
   }
   id, err := res.(sql.Result).LastInsertId()
   if err != nil {
      return "", err
   }
   return strconv.FormatInt(id, 10), nil
}

// 執行SQL指令
func (m *MySQL) Exec(sqlstr string, cond ...interface{}) (interface{}, error) {
   m.CheckAndReConnect()
   stmt, err := m.Conn.Prepare(sqlstr)
   if err != nil {
      return nil, fmt.Errorf("Prepare SQL error: %v", err)
   }
   res, err := stmt.Exec(cond...)
   defer stmt.Close()
   if err != nil {
      return nil, err
   }
   id, err := res.(sql.Result).LastInsertId()
   if err != nil {
      return nil, err
   }
   return id, nil
}

// 判斷資料是否存在 true)存在 false)不存在
func (m *MySQL) RowExists(sqlstr string, args ...interface{}) bool {
   m.CheckAndReConnect()
   var exists bool = true // interface{}
   s := fmt.Sprintf("select exists(%s)", sqlstr)
   row := m.Conn.QueryRow(s, args...)
   err := row.Scan(&exists)
   if err != nil {
      if err == sql.ErrNoRows || err.Error() == "sql: no rows in result set" {  // 不存在
         exists = false
      } else {
         exists = true
      }
   } 
   return exists
}

// bl, err := conn.RowExistsAndError(sql, user.Name)
func(m *MySQL) RowExistsAndError(sql string, args ...interface{})(bool, error) {
   m.CheckAndReConnect()
   rows, err := m.Conn.Query(sql, args...)
   if err != nil {
      if err.Error() == "sql: no rows in result set" {
         return false, nil
      } else {
         return false, err
      }
   }
   count := 0
   for rows.Next() {
      _ = rows.Scan(&count)
   }   
   if count > 0 {
      return true, nil
   } else {
      return false, nil
   }
}

// 結束資料庫連線
func (m *MySQL) Disconnect() {
   defer m.Conn.Close()
}

// 建立資料庫連線 aka doreconnect()
func (m *MySQL) Database(config DBConnect) (error) {
   if len(config.DBMS) == 0 {
      config.DBMS = "mysql"
   }
   if len(config.DbPort) == 0 {
      config.DbPort = "3306"
   }
   if config.DbServer == "" || config.DbName == "" || config.DbLogin == "" || config.DbPasswd == "" {
      return fmt.Errorf("Db Server setup params is wrong.")
   }

   var b bytes.Buffer
   b.WriteString(config.DbLogin)
   b.WriteString(":")
   b.WriteString(config.DbPasswd)
   b.WriteString("@tcp(")
   b.WriteString(config.DbServer)
   b.WriteString(":")
   b.WriteString(config.DbPort)
   b.WriteString(")/")
   b.WriteString(config.DbName)
   b.WriteString("?charset=utf8")

   conn, err := sql.Open(config.DBMS, b.String())
   if err != nil {
       return err
   }
   m.Conn = conn 
   m.Config = config
   return nil
}

// 建立資料庫連線
func NewSherryDB(config DBConnect) (*MySQL, error)  {
   mysql := MySQL{}

   if err := mysql.Database(config); err != nil {
      return nil, err
   }

   return &mysql, nil
}
