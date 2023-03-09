Package SherryDB

import(
   "fmt"
   "gorm.io/gorm"
   "gorm.io/driver/sqlite"
)

type SQLite struct {
   Conn		*gorm.DB	// Database connection
   FileName	string		// 檔案名稱
}

func NewSQLite(filePathAndName string)(*SQLite, error) {
   if filePathAndName == "" {
      return nil, fmt.Errorf("file path and file name is empty")
   }
   db, err := gorm.Open(sqlite.Open(filePathAndName), &gorm.Config{})
   if err != nil {
      return nil, err
   }
   return &SQLite {
      Conn: db,
      FileName: filePathAndName,
   }, nil
}
