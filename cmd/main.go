package main

import (
	"fmt"
	"go.zakaria.org/crtsh"
	"log"
)

func main() {
	r, err := crtsh.SearchJSON("bsd.lv", false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", r)
	//for _, s := range r {
	//	fmt.Printf("==>\n")
	//	fmt.Printf("%+v\n", s)
	//}
}
