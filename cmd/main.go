// Package main creates a command that gets konnakol line from stdin and prints
// formatted string to stdout
package main

import (
	"flag"
	"log"
	"os"

	konnakolliner "github.com/asahnoln/konnakol-liner"
)

func main() {
	thalam := flag.Int("t", konnakolliner.ThalamAdi, "set thalam count")
	gathi := flag.Int("g", konnakolliner.GathiChatushram, "set gathi count")
	flag.Parse()

	log.Fatal(konnakolliner.LineOut(os.Stdin, os.Stdout, *thalam, *gathi))
}
