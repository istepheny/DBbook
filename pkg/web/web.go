package web

import (
	"dbbook/pkg/config"
	"log"
	"net/http"
	"path/filepath"
)

func Serve(s config.Server) {

	appPath, _ := filepath.Abs(".")

	fileServer := http.FileServer(http.Dir(appPath + "/web"))

	log.Printf("DBbook is running at: http://%s:%s\n", s.Host, s.Port)

	e := http.ListenAndServe(s.Host+":"+s.Port, fileServer)

	if e != nil {
		log.Fatalf("Failed to start service: %s", e)
	}
}
