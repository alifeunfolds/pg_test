package main

import (
  "database/sql"
  _ "github.com/lib/lq"

  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

func main() {
  db, err := sql.Open("","")

  e := echo.New()
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.Logger.Fatal(e.Start(":8000"))
}
