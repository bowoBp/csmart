package main

import (
	"fmt"
	"github.com/bowoBp/csmart/modules/api"
)

func main() {
	//router := gin.Default()
	app := api.Default()
	if err := app.Start(); err != nil {
		panic(fmt.Errorf("ap.start: %w", err))
	}

	//onlineBook := bookingbook.NewRequestHandler()
	//onlineBook.Handle(router)

}
