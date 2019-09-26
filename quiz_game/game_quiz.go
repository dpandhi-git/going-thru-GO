package main

import (
	"fmt"
	"os"
	"flag"
	"encoding/csv"
)

func main(){
	csvFilepath := flag.String("csv","problems.csv","a csv file format with 'question/answers")
	flag.Parse()

	file,err := os.Open(*csvFilepath)
	if err != nil {
		exit(fmt.Sprintf("Failed to load the file": %s \n", *csvFilepath))
	}

}