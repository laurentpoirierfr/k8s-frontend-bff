package main

import "gofr.dev/pkg/gofr"

func main() {
	// initialise gofr object
	app := gofr.New()

	// register route service-02
	app.GET("/api/service-02", func(ctx *gofr.Context) (interface{}, error) {
		return "service-02", nil
	})
	
	// Runs the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Run()
}
