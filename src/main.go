package main

import (
	"fmt"

	m3uparser "github.com/pawanpaudel93/go-m3u-parser/m3uparser"
)

func main() {
	parser := m3uparser.M3uParser{}
	parser.ParseM3u("https://iptv-org.github.io/iptv/index.m3u", true, true)
	parser.FilterBy("status", []string{"GOOD"}, true)
	parser.SortBy("category", true)
	fmt.Println("Saved stream information: ", len(parser.GetStreamsSlice()))
	parser.ToFile("listaTeste.json")
}
