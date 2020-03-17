package helper

import (
	"dbbook/pkg/flags"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func AppPath() string {
	appPath, _ := filepath.Abs(".")

	return appPath
}

func ConfigFilePath() string {
	return strings.Join([]string{AppPath(), "database.json"}, string(os.PathSeparator))
}

func TemplatePath() string {
	return strings.Join([]string{AppPath(), "web", "template", ""}, string(os.PathSeparator))
}

func BookPath() string {
	return strings.Join([]string{AppPath(), "web", "dbbook", ""}, string(os.PathSeparator))
}

func Mkdir(dir string) {
	if _, e := os.Stat(dir); os.IsNotExist(e) {
		e := os.MkdirAll(dir, 0755)
		if e != nil {
			log.Fatal(e)
		}
	}
}

type IP struct {
	Local   string
	Network string
}

func GetIp() (ip IP) {
	addrs, e := net.InterfaceAddrs()
	if e != nil {
		log.Fatalf("Failed to get ip: %s", e)
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.To4() != nil {
			switch ipnet.IP.IsLoopback() {
			case true:
				ip.Local = ipnet.IP.String()
			case false:
				if !strings.HasPrefix(ipnet.IP.String(), "169.254.") {
					ip.Network = ipnet.IP.String()
				}
			}
		}
	}

	return ip
}

func RunningLog() {
	ip := GetIp()

	format := "DBbook is running at:\n" +
		"- Local:	http://%s:%s\n"

	if ip.Network != "" {
		format += "- Network:	http://%s:%s"
		log.Printf(format, ip.Local, flags.Port, ip.Network, flags.Port)
	} else {
		log.Printf(format, ip.Local, flags.Port)
	}
}
