package main

import (
  "fmt"
  "net/http"
  //"time"

  "database/sql"
  _ "github.com/lib/pq"

  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

const (
    DB_USER = "postgres"
    DB_PASSWORD = "postgres"
    DB_NAME = "test"
)

func main() {

  e := echo.New()
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // database connection
  dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
          DB_USER, DB_PASSWORD, DB_NAME)
  db, err := sql.Open("postgres",dbinfo)
  if err != nil {
    e.Logger.Fatal(err)
  }
  defer db.Close()

  // Create
  e.POST("/users", func(c echo.Context) error {
    var lastInsertId int
    err = db.QueryRow("INSERT INTO userinfo(username,departname, created) VALUES($1,$2,$3) returning uid;", "alifeunfolds", "safecorners", "2017-06-06").Scan($lastInsertId)
    if err != nil {
      return c.JSON(403,"Internal Server Error")
    }
    return C.JSON(HTTP.StatusOK, "lastInsertId: %d", lastInsertId)
  })

  // Read
  e.GET("/users/:id", func(c echo.Context) error {
    fmt.Println("# Querying")
    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)

    for rows.Next() {
        var uid int
        var username string
        var department string
        var created time.Time
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println("uid | username | department | created ")
        fmt.Printf("%3v | %8v | %6v | %6v\n", uid, username, department, created)
    }
  })
  // Update
  e.PUT("/users/:id", func(c echo.Context) error {
    fmt.Println("# Updating")
    stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
    checkErr(err)

    res, err := stmt.Exec("astaxieupdate", lastInsertId)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect, "rows changed")
  })

  // Delete
  e.DELETE("/users/id", func(c echo.Context) error {

    stmt, err := db.Prepare("delete from userinfo where uid=$1")
    checkErr(err)

    res, err = stmt.Exec(lastInsertId)
    checkErr(err)

    affect, err = res.RowsAffected()
    checkErr(err)

    fmt.Println(affect, "rows changed")
  })

  e.Logger.Fatal(e.Start(":8000"))
}
