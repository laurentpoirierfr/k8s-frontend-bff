package main

import (
	"encoding/json"
	"io"
	"os"

	"gofr.dev/pkg/gofr"
)

type Result struct {
	Data string `json:"data"`
}

func main() {
	// initialise gofr object
	app := gofr.New()

	// register a payment service which is hosted at http://localhost:9000
	app.AddHTTPService("service-01", os.Getenv("HTTP_URL_SERVICE_01"))

	// register route service-01
	app.GET("/api/service-01", func(ctx *gofr.Context) (interface{}, error) {

		// Get the payment service client
		service := ctx.GetHTTPService("service-01")

		// Use the Get method to call the GET /user endpoint of payments service
		resp, err := service.Get(ctx, "/api/service-01", nil)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var result Result
		err = json.Unmarshal(body, &result)

		return result.Data, err
	})

	// register a payment service which is hosted at http://localhost:9000
	app.AddHTTPService("service-02", os.Getenv("HTTP_URL_SERVICE_02"))

	// register route service-01
	app.GET("/api/service-02", func(ctx *gofr.Context) (interface{}, error) {

		// Get the payment service client
		service := ctx.GetHTTPService("service-02")

		// Use the Get method to call the GET /user endpoint of payments service
		resp, err := service.Get(ctx, "/api/service-02", nil)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var result Result
		err = json.Unmarshal(body, &result)

		return result.Data, err
	})

	// Runs the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Run()
}
