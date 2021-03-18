package main

import (
	"dbbook/pkg/config"
	"dbbook/pkg/document"
	"dbbook/pkg/schema"
	"dbbook/pkg/web"
	"time"
)

func main() {
	c := config.Load()

	go func() {
		ticker := time.NewTicker(time.Second * time.Duration(c.Ticker))
		for ; true; <-ticker.C {
			book := schema.Query(c.Databases)
			document.Write(book)
		}
	}()

	web.Serve(c.Server)
}
