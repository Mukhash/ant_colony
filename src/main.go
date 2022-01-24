package main

import (
	"fmt"
	lemin "lem-in/lemin"
	"log"
	"os"
)

func main() {
	arg := os.Args
	if len(arg) != 2 {
		log.Fatal("invalid number of arguments")
	}
	colonyFile, err := os.ReadFile(arg[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(colonyFile))
	fmt.Println()
	if !lemin.ValidateFile(colonyFile) {
		{
			log.Fatal("File invalid")
		}
	}
	colony := lemin.Colony{}
	colony.Init()
	err = lemin.PopulateColony(colonyFile, &colony)
	if err != nil {
		log.Fatal(err)
	}
	if err = lemin.CheckColony(&colony); err != nil {
		log.Fatal(err)
	}
	lemin.Lemin(&colony)
}
