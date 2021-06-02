package main

import (
	"flag"
	"fmt"
	"github.com/gusanmaz/usableurl"
)

var (
	flagURLValue string
	flagURLUsage string = "URL that may need processing (expansion, addition of http(s) prefix etc..)"
)

func init() {
	flag.StringVar(&flagURLValue, "url", "", flagURLUsage)
	flag.StringVar(&flagURLValue, "u", "", flagURLUsage + " (shorthand)")
	flag.Parse()
}

func main() {
	fmt.Printf("Input URL: %v\n", flagURLValue)
	fmt.Printf("Output URL: %v\n", usableurl.Sanitize(flagURLValue))
}
