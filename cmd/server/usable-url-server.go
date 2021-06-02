package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gusanmaz/usableurl"
	"log"
	"net/http"
	"strings"
)

type Data struct {
	In  string
	Out string
}

var (
	flagPortValue   string
	flagPortUsage   string = "Port for web server"
	flagPortDefault string = "9001"
)

func init() {
	flag.StringVar(&flagPortValue, "url", flagPortDefault, flagPortUsage)
	flag.StringVar(&flagPortValue, "u", flagPortDefault, flagPortUsage+"(shorthand)")
	flag.Parse()
}

var urlHandler = func(w http.ResponseWriter, r *http.Request) {
	urlIn := r.URL.Query().Get("url")
	urlIn = strings.TrimSpace(urlIn)
	urlOut := usableurl.Sanitize(urlIn)
	urlOut = strings.TrimSpace(urlOut)
	d := Data{
		In:  urlIn,
		Out: urlOut,
	}
	bytes, _ := json.Marshal(d)

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func main() {
	portAddr := fmt.Sprintf("localhost:%v", flagPortValue)
	http.HandleFunc("/", urlHandler)
	fmt.Println(portAddr)

	err := http.ListenAndServe(portAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
