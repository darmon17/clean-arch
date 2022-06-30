package main

import (
	"project2/config"
	"project2/factory"
	"project2/migration"
	"project2/routes"
)

// go get -u github.com/labstack/echo/v4 (echo)
// go get github.com/jinzhu/now (gorm)
// go get github.com/golang-jwt/jwt (authentication)
// go get -u gorm.io/driver/mysql (gorm mysql)
// go get github.com/labstack/echo/v4/middleware panggil jwt

func main() {
	dbConn := config.InitDB()
	migration.InitMigrate(dbConn)
	presenter := factory.InitFactory(dbConn)

	e := routes.New(presenter)

	e.Logger.Fatal(e.Start(":80"))
}
