package main

import (
	"dbbook/pkg/config"
	"dbbook/pkg/document"
	"dbbook/pkg/flags"
	_ "dbbook/pkg/flags"
	"dbbook/pkg/schema"
	"dbbook/pkg/web"
	"time"
)

func main() {
	databases := config.Load()

	go func() {
		ticker := time.NewTicker(time.Second * time.Duration(flags.Ticker))
		for ; true; <-ticker.C {
			book := schema.Query(databases)
			document.Write(book)
		}
	}()

	web.Serve()
}
