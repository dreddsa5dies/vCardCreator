package main

import (
	"bytes"
	"fmt"
	"image/png"
	"io"
	"log"
	"os"

	"flag"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %v -h\n", os.Args[0])
		os.Exit(1)
	} else {
		// parse flags
		fVcardOpen := flag.String("o", "", "Open vCard file (all format)")
		fBarcodeSave := flag.String("s", "", "Get name for save barcode file")

		flag.Parse()

		// need read file
		// read file as a string
		fOpen, err := os.Open(*fVcardOpen)
		if err != nil {
			log.Fatalf("File open ERROR: %v\n", err)
		}
		defer fOpen.Close()
		buf := bytes.NewBuffer(nil)
		io.Copy(buf, fOpen)

		// Create the barcode from data on file open
		qrCode, err := qr.Encode(string(buf.Bytes()), qr.M, qr.Auto)
		if err != nil {
			log.Fatalf("Create the barcode ERROR: %v\n", err)
		}

		// Scale the barcode to 200x200 pixels
		qrCode, err = barcode.Scale(qrCode, 200, 200)
		if err != nil {
			log.Fatalf("Scale the barcode ERROR: %v\n", err)
		}

		// create the output file
		// format png
		file, err := os.Create(*fBarcodeSave + ".png")
		if err != nil {
			log.Fatalf("File save ERROR: %v\n", err)
		}
		defer file.Close()

		// encode the barcode as png
		png.Encode(file, qrCode)
	}
}
