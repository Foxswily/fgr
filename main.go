package main

import (
	"flag"
	"github.com/Foxswily/fgr/template"
	"log"
)

func main() {
	opData := flag.String("d", "in.yaml", "data file")
	flag.Parse()

	data, err := template.Read(*opData)
	if err != nil {
		log.Printf("read %s error %s\n", *opData, err)
	}
	err = template.Write(data)
	if err != nil {
		log.Printf("write %s error %s\n", *opData, err)
	}
}
