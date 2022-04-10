package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
	"github.com/liran/http-server/ip"
)

var dirHandle = http.FileServer(http.Dir("."))

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		return
	}

	dirHandle.ServeHTTP(w, req)
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// 参数处理
	port := flag.Int("p", 8080, "Port")
	cors := flag.Bool("cors", true, "Enable CORS via the \"Access-Control-Allow-Origin\" header")
	openbrowser := flag.Bool("o", false, "Open browser window after starting the server")
	flag.Parse()

	// 默认支持跨域
	if *cors {
		http.HandleFunc("/", indexHandler)
	} else {
		http.Handle("/", dirHandle)
	}

	portString := fmt.Sprintf(":%d", *port)
	if *openbrowser {
		openBrowser("http://localhost" + portString)
	}

	color.Cyan("Starting up http-server...")
	// 显示可用的IP
	ip.ShowAvailableIps(*port)
	color.Red("Hit CTRL-C to stop the server")

	log.Fatal(http.ListenAndServe(portString, nil))
}
