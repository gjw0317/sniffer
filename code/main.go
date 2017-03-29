package main

import (
	"Sniffer/code/rule"
	"Sniffer/code/sniffer"
	"flag"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
)

var (
	path  = flag.String("path", "../images", "path holds the path to store the pictures")
	start = flag.Int("start", 1, "start holds the first page to start sniffer")
	end   = flag.Int("end", 1, "end holds the last page number to finish sniffer")
)

var client *http.Client = &http.Client{}

func init() {
	flag.Parse()
	// check parameters
	if *start > *end {
		fmt.Println("Invalid Arguments\t:--start(" + strconv.Itoa(*start) + ") > end(" + strconv.Itoa(*end) + ")")
	}

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	sniffer.New(*path, *start, *end, rule.NewJandanRule(), client).Run()
}
