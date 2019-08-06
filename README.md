## 資料庫連線工具 SherryDB

### How to use it ?
* install
```
go get github.com/asccclass/sherrydb/mysql
```

* example
```
package main

import (
   "fmt"
   "github.com/asccclass/sherrydb/mysql"
   "os"
)

var dbconnect sherrydb.DBConnect

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
