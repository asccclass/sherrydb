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

   conn, err := SherryDB.NewSherryDB(dbconnect)
   if err != nil {
      fmt.Printf("%v", err)
      return 
   } 
   orgs, err := conn.DoreFetchHash("select * from organizations")
   if err != nil {
      fmt.Printf("%v", err)
      return
   } 
   fmt.Printf("%v\n", orgs)
/*
   for _, d := range orgs {
      fmt.Printf("%v\n", string(d))
   }
*/
   conn.Disconnect() 
}
