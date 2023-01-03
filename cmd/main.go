// Package main creates a command that gets konnakol line from stdin and prints
// formatted string to stdout
package main

import (
	"flag"
	"log"
	"os"

	liner "github.com/asahnoln/konnakol-liner"
)

func main() {
	thalam := flag.Int("t", liner.ThalamAdi, "set thalam count")
	gathi := flag.Int("g", liner.GathiChatushram, "set gathi count")
	flag.Parse()

	log.Fatal(liner.Out(os.Stdin, os.Stdout, liner.Highlight, *thalam, *gathi))
}
