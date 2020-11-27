package server

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"tiket.vip/src/infrastructures/configs"
	elastics "tiket.vip/src/infrastructures/elastic"
	loggers "tiket.vip/src/infrastructures/logger"
	"tiket.vip/src/infrastructures/middlewares"
	orm "tiket.vip/src/infrastructures/orm/gorm"
	routers "tiket.vip/src/interfaces/routes/v1"
)

func Serve() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	datasources := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))
	db, errDb := sql.Open("mysql", datasources)
	if errDb != nil {
		log.Fatal(fmt.Sprintf("Error connecting db, %s", errDb.Error()))
	}
	db.SetMaxIdleConns(10)

	db.SetMaxOpenConns(100)

	db.SetConnMaxLifetime(time.Hour)

	orm, errOrm := orm.New(db)
	if errDb != nil {
		log.Panic(fmt.Sprintf("Error connecting db, %s", errOrm.Error()))
	}
	loggers.Init()
	e := New()

	authMiddle := middlewares.NewAuthMiddleware(orm)

	h := routers.NewHandler(orm, elastics.Init())
	e.Use(authMiddle.MiddleGate())
	v1 := e.Group("/api")
	h.Register(v1)

	appHost := os.Getenv("APPLICATION_PORT")

	if "" == appHost {
		log.Panic("key of APPLICATION HOST are not define.")
	}

	e.Logger.Fatal(e.Start(":" + appHost))

}

func New() *echo.Echo {
	eHandler := middlewares.NewErrorMiddleware()

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.Validator = configs.NewValidator()
	e.HTTPErrorHandler = eHandler.HttpErrorHandler

	return e
}
