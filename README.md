# SherryDB library

SherryDB is a lightweight database connection utility written in Go.  
It simplifies connecting to databases like MySQL and SQLite using environment variables, making it ideal for small projects, prototyping, or educational purposes.

## Features

- Supports multiple database management systems (DBMS), including MySQL and SQLite.
- Utilizes environment variables for configuration, enhancing flexibility and security.
- Provides a straightforward API for establishing and managing database connections.
- Includes example code to help you get started quickly.

## Installation

To install SherryDB, use `go get` to fetch the necessary packages:

```bash
go get github.com/go-sql-driver/mysql
go get github.com/asccclass/sherrydb/mysql
```

Alternatively, if you're using Go modules, initialize your module:

```bash
go mod init your_module_name
```

## Usage

### Setting Up Environment Variables

Before running your application, set the following environment variables to configure your database connection:

- `DBMS`: The type of your database management system (e.g., `mysql` or `sqlite`).
- `DBSERVER`: The address of your database server (e.g., `localhost`).
- `DBPORT`: The port number your database server is listening on (e.g., `3306`).
- `DBNAME`: The name of your database.
- `DBLOGIN`: Your database username.
- `DBPASSWORD`: Your database password.

You can set these variables in your shell or define them in an `.env` file.

### Example Code

Here's a simple example demonstrating how to use SherryDB to connect to a database:

```go
package main

import (
    "fmt"
    "os"

    "github.com/asccclass/sherrydb/mysql"
)

func main() {
    dbconnect := mysql.DBConnect{
        DBMS:     os.Getenv("DBMS"),
        DbServer: os.Getenv("DBSERVER"),
        DbPort:   os.Getenv("DBPORT"),
        DbName:   os.Getenv("DBNAME"),
        DbLogin:  os.Getenv("DBLOGIN"),
        DbPasswd: os.Getenv("DBPASSWORD"),
    }

    conn, err := mysql.NewSherryDB(dbconnect)
    if err != nil {
        fmt.Printf("Connection error: %v\n", err)
        return
    }
    defer conn.Conn.Close()

    fmt.Printf("Connected successfully: %v\n", conn)
}
```

This code initializes a new database connection using the parameters provided via environment variables.

## Project Structure

- `mysql/`: Contains the MySQL-specific implementation of the SherryDB connector.
- `sqlite/`: Contains the SQLite-specific implementation of the SherryDB connector.
- `example.go`: Provides an example of how to use SherryDB in your application.
- `envfile`: A sample file demonstrating how to set environment variables.
- `go.mod` and `go.sum`: Go module files managing dependencies.
- `makefile`: Automates build and setup tasks.

## Contributing

Contributions are welcome! If you have suggestions for improvements or have found bugs, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
