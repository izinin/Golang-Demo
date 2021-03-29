package main

import (
	"fmt"
	"net/http"

	"github.com/buger/jsonparser"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const port = ":4000"

// https://catalog.data.gov/dataset/househartford-comittment-2005-to-june-1-2014
const data_source = "https://data.hartford.gov/api/views/62ub-3292/rows.json?accessType=DOWNLOAD"

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", get_data)

	// Start server
	e.Logger.Fatal(e.Start(port))
}

// Handler
func get_data(c echo.Context) error {
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

	items := make([]HousingPrice, 0, 50)
	jsonparser.ArrayEach(resp.Body(), func(record []byte, dataType jsonparser.ValueType, offset int, err error) {
		items = append(items, *newHousingPrice(record))
	}, "data")
	fmt.Println(items)

	return c.String(http.StatusOK, resp.String())
}
