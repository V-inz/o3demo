// This program uses github.com/skip2/go-qrcode to generate the QR-Code you need,
// to get Threema-Status green for your bots.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/o3ma/o3"

	// https://github.com/skip2/go-qrcode
	// $ go get -u github.com/skip2/go-qrcode/...
	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	var (
		pass = []byte{0xA, 0xB, 0xC, 0xD, 0xE}
		idpath  = "threema.id"
		tid     o3.ThreemaID
	)

	// check whether an id file exists
	if _, err := os.Stat(idpath); err != nil {
		fmt.Printf("Identity-file '%s' missing.\n", idpath)
	} else {
		// load existing ID
		tid, err = o3.LoadIDFromFile(idpath, pass)
		if err != nil {
			log.Fatal(err)
		}
	}

	// concat QR-Code content. "3mid" maybe is a shortcut for Threema-ID
	qrtext := fmt.Sprintf("3mid:%s,%x", tid.String(), tid.GetPubKey()[:])
	fmt.Printf(qrtext + "\n")

	// generate the PNG-Image "threemaid.png"
	err := qrcode.WriteFile(qrtext, qrcode.Medium, 256, "threemaid.png")
	if err != nil { 
		log.Fatal(err)
	}
}
