package SherryDB

import (
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "fmt"
)

// 資料庫連線設定
type DBConnect struct {
   DBMS     string
   DbServer string
   DbPort   string
   DbName   string
   DbLogin  string
   DbPasswd string
}

type MySQL struct {
   Name string
   Conn *sql.DB
}

func (m *MySQL) Disconnect() {
   defer m.Conn.Close()
}

func (m *MySQL) DoreFetchHash(sqlString string) (string, error) {
   tableData := make([]map[string]interface{}, 0)

   rows, err := m.Conn.Query(sqlString)
   if err != nil {
      return "", err
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

func NewSherryDB(config DBConnect) (*MySQL, error)  {
   if len(config.DBMS) == 0 {
      config.DBMS = "mysql"
   }
   if len(config.DbPort) == 0 {
      config.DbPort = "3306"
   }

   if config.DbServer == "" || config.DbName == "" || config.DbLogin == "" || config.DbPasswd == "" {
      return nil, fmt.Errorf("Db Server setup params is wrong.")
   }

   conn, err := sql.Open(config.DBMS, config.DbLogin + ":" + config.DbPasswd + "@tcp("+config.DbServer+":" + config.DbPort+")/"+config.DbName+"?charset=utf8")
   if err != nil {
      return nil, err
   }
   return &MySQL {
      Name: config.DBMS,
      Conn: conn,
   }, nil
}
