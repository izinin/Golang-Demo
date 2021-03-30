package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/buger/jsonparser"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const port = ":4200"

// https://catalog.data.gov/dataset/househartford-comittment-2005-to-june-1-2014
const data_source = "https://data.hartford.gov/api/views/62ub-3292/rows.json?accessType=DOWNLOAD"
const dsn = "host=db user=dbuser password=pass dbname=demo port=5432 sslmode=disable"

type CustomContext struct {
	echo.Context
	db *gorm.DB
}

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				log.Fatal("cannot connect to dbase: %s", err.Error())
			}
			cc := &CustomContext{c, db}
			return next(cc)
		}
	})
	// Routes
	e.GET("/import", import_db)

	// Start server
	e.Logger.Fatal(e.Start(port))
}

// Handler
func import_db(c echo.Context) error {
	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get(data_source)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot get data from %s :  %s", data_source, err.Error())
	}

	// get data title
	descr, err := jsonparser.GetString(resp.Body(), "meta", "view", "description")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot parse data for path ")
	}

	fmt.Println(descr)

	cc := c.(*CustomContext)
	jsonparser.ArrayEach(resp.Body(), func(record []byte, dataType jsonparser.ValueType, offset int, err error) {
		result := cc.db.Create(newHousingPrice(record))
		fmt.Println(result)
	}, "data")

	return c.String(http.StatusOK, resp.String())
}
