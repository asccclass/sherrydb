## 資料庫連線工具 SherryDB

### How to use it ?
* install
```
go get github.com/go-sql-driver/mysql
go get github.com/asccclass/sherrydb/mysql
```

### Usage
* Connect database
```
package main

import (
   "fmt"
   "github.com/asccclass/sherrydb/mysql"
   "os"
)

var dbconnect SherryDB.DBConnect

func main() {
   // check DB Information
   dbconnect.DBMS = os.Getenv("DBMS")
   dbconnect.DbServer = os.Getenv("DBSERVER")
   dbconnect.DbPort = os.Getenv("DBPORT")
   dbconnect.DbName = os.Getenv("DBNAME")
   dbconnect.DbLogin = os.Getenv("DBLOGIN")
   dbconnect.DbPasswd = os.Getenv("DBPASSWORD")

   conn, err := sherrydb.NewSherryDB(dbconnect)
   defer conn.Conn.Close()
   if err != nil {
      fmt.Printf("%v", err)
   } else {
      fmt.Printf("Connect ok %v", conn)
   }
}
```

### Create Record

```
sql := "insert into beacon(ip,mac,beaconID,lastupdate) value(?,?,?,?)"
_, err = conn.Conn.Exec(sql, ip.LocalIP,ip.MACAdress,beaconID,st.Now())
```

### Read record

```
row := conn.Conn.QueryRow("select ip,mac,beaconID,lastupdate from beacon where mac=?", ip.MACAddress)
var ip,mac,beaconID,lastupdate string
st := sherrytime.NewSherryTime("Asia/Taipei", "-")  // Initial
if err := row.Scan(&ip,&mac,&beaconID,&lastupdate); err != nil {
   switch {
      case err == sql.ErrNoRows:   // No data. Since you use sql.ErrNoRows, you need import "database/sql"
         // do your code
         return
   }
   fmt.Println(err.Error())
   return
}
```

### Update Record

```
sql := "update beacon set ip=?,mac=?,beaconID=?,lastupdate=? where mac=?"
conn.Conn.Exec(sql, ip.LocalIP,ip.MACAdress,beaconID,st.Now(),ip.MACAddress)
```

* DoreSelOne
```
X := Ranks{}
val, err := conn.DoreSelOne("select * from rank where rankID=?", &X, rankID)
if err != nil {
   w.WriteHeader(http.StatusInternalServerError)
   fmt.Fprintf(w, "{\"errMsg\": \"Record not found.\"}")
   return
}
```

* Get Insert's auto imcrement key
```
id, err := res.(sql.Result).LastInsertId()
```
