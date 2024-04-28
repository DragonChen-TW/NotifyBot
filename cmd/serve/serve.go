package main

import (
	"github.com/dragonchen-tw/NotifyBot/pkg/notify"
)

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Serve
// Date:	2024/04/24

// func main() {
// 	e := echo.New()
// 	e.GET("/", Hello)

// 	// Start server
// 	e.Logger.Fatal(e.Start(":8080"))
// }

// // Handler
// func Hello(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World!")
// }

func main() {
	notify.LINE()
}
