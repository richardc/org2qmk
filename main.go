package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/niklasfasching/go-org/org"
	"github.com/richardc/org2qmk/org2qmk"
)

func main() {
	log := log.New(os.Stderr, "", 0)
	if len(os.Args) < 3 {
		log.Println("USAGE: org2qmk FILE OUTPUT_FORMAT")
		log.Fatal("Supported output formats: keymap")
	}
	path := os.Args[1]
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	out, err := "", nil
	d := org.New().Parse(bytes.NewReader(bs), path)
	switch strings.ToLower(os.Args[2]) {
	case "keymap":
		out, err = d.Write(org2qmk.NewQmkKeymapWriter())
	default:
		log.Fatal("Unsupported output format")
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(os.Stdout, out)
}
