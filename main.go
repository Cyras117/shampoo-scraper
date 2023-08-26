package main

import (
	"fmt"
	"os"
	"shampoo-scraper/src/utils"
	"strings"
)

func main() {
	args := os.Args[1:]
	if args == nil {
		fmt.Println("Error:Args missing")
	}

	for iterator, arg := range args {
		//TODO:Ccreate a function to handdle this
		if !utils.IsIn("asura", strings.ToLower(arg)) {
			fmt.Printf("Error:Argument %d \"%s\" does not have an origin site paramiter\n", iterator, arg)
		}
	}
}
