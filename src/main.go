package main

import (
	"bytes"
	"image/png"
	"io"
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

var opts struct {
	fVcardOpen   string `short:"o" long:"open" default:"vCard.txt" description:"open vCard file"`
	fBarcodeSave string `short:"s" long:"save" default:"vCard.png" description:"save barcode file"`
}

func main() {
	// parse flags
	flags.Parse(&opts)

	// need read file
	fOpen, err := os.Open(opts.fVcardOpen)
	if err != nil {
		log.Printf("File open ERROR: %v\n", err)
	}
	defer fOpen.Close()
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, fOpen)

	// Create the barcode from data on file open
	qrCode, err := qr.Encode(string(buf.Bytes()), qr.M, qr.Auto)
	if err != nil {
		log.Printf("Create the barcode ERROR: %v\n", err)
	}

	// Scale the barcode to 200x200 pixels
	qrCode, err = barcode.Scale(qrCode, 200, 200)
	if err != nil {
		log.Printf("Scale the barcode ERROR: %v\n", err)
	}

	// create the output file
	file, err := os.Create(opts.fBarcodeSave)
	if err != nil {
		log.Printf("File save ERROR: %v\n", err)
	}
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
}
