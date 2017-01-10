// This little program uses qrencode to generate the QR-Code you need
// to get Threema-Status "green" for your bots.
// You need Package qrencode installed on your system:  # apt-get install qrencode
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/o3ma/o3"
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
	qrcode := fmt.Sprintf("3mid:%s,%x", tid.String(), tid.GetPubKey()[:])
	fmt.Printf(qrcode + "\n")

	// shell-execute command to generate the PNG-Image "threemaid.png"
	qrcode_command := fmt.Sprintf("qrencode  -l Q -o threemaid.png \"%s\"", qrcode)
	c := exec.Command("sh","-c", qrcode_command)
	if err := c.Run(); err != nil { 
		fmt.Println("Error: ", err)
	}
}
