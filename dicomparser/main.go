package main

import (
	"flag"
	"fmt"
	"github.com/mpetavy/go-dicom-parser/dicom"
	"log"
	"os"
)

var fn *string

func main() {
	fn = flag.String("file","","dicom filename to parse")
	flag.Parse()

	if *fn == "" {
		flag.Usage()
		os.Exit(1)
	}

	r, err := os.Open(*fn)
	if err != nil {
		log.Fatalf("os.Open(_) => %v", err)
	}
	dataSet, err := dicom.Parse(r)
	if err != nil {
		log.Fatalf("dicom.Parse(_) => %v", err)
	}

	for tag, element := range dataSet.Elements {
		fmt.Println(tag, element.Tag,dicom.GetName(tag),element.VR, element.ValueField)
	}
}