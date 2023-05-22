package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"go.zakaria.org/crtsh"
)

var (
	jsonOutput bool
	expired    bool
	query      string
)

func usage() {
	fmt.Printf(`usage: crtsh [-e] [-j] query
where:
	-e	show expired certs
	-j	output raw JSON returned
	query	the domain query
`)
}

func init() {
	flag.Usage = usage
	flag.BoolVar(&expired, "e", false, "show expired certs")
	flag.BoolVar(&jsonOutput, "j", false, "output raw json instead of formatted output")
	flag.Parse()

	query = flag.Arg(0)
}

func main() {

	if len(query) == 0 {
		log.Print("no query given")
		usage()
		os.Exit(1)
	}

	if jsonOutput {
		r, err := crtsh.SearchJSON(query, expired)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", r)
	} else {
		r, err := crtsh.Search(query, expired)
		if err != nil {
			log.Fatal(err)
		}

		for _, e := range r {
			fmt.Printf("==> %d (%s) [%s -> %s]\n", e.Id, e.EntryTimestamp, e.NotBefore, e.NotAfter)
			fmt.Printf("\tIssuer: %s (%d)\n", e.IssuerName, e.IssuerCaId)
			fmt.Printf("\tCommon Name: %s\n", e.CommonName)
			names := strings.ReplaceAll(e.NameValue, "\n", ",")
			fmt.Printf("\tNames: %s\n", names)
		}

	}
}
