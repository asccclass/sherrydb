package main

import (
   "github.com/asccclass/sherrydb/mysql"
   "fmt"
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

   dbconnect.DBMS = os.Getenv("DBMS")
   dbconnect.DbServer = "192.168.12.223"
   dbconnect.DbPort = "3306"
   dbconnect.DbName = "learnpoints"
   dbconnect.DbLogin = "ascal"
   dbconnect.DbPasswd = "webteam@2018"

   conn, err := SherryDB.NewSherryDB(dbconnect)
   if err != nil {
      fmt.Printf("%v", err)
      return 
   } 
   var sysid string = ""
   var course string = "(201919)資通安全管理法及維護計畫簡介說明會(第二場)"
   var y int = 2019
   if !conn.RowExists("select * from excela where Sysid=? and Year=? and CourseTitle=?", sysid, y, course) {
      fmt.Println("No data.")
   }
   _, err = conn.DoreFetchHash("select * from excela")
   if err != nil {
      fmt.Printf("%v", err)
      return
   } 
   // fmt.Printf("%v\n", orgs)
/*
   for _, d := range orgs {
      fmt.Printf("%v\n", string(d))
   }
*/
   conn.Disconnect() 
}
