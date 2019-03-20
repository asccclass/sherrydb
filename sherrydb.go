package main

import (
   "./mysql"
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
   } else {
      fmt.Printf("Connect ok %v", conn)
   }
   conn.Disconnect() 
}
