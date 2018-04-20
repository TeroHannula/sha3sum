// sha3sum hash calculator for 224, 256, 384 and 512 bit hashes
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/sha3"
)

func main() {

	var hashlengthPtr *int = flag.Int("l", 512, "hash length in bits, use values 224, 256, 384 or 512")
	var quiet *bool = flag.Bool("q", false, "Do not print filename and hash length in the output")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Filename missing")
		return
	}

	filename := flag.Args()[0]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("File %s not found\n", filename)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("File read error: %s\n", err)
		return
	}

	success := false
	switch *hashlengthPtr {
	case 224:
		fmt.Printf("%x\n", sha3.Sum224(data))
		success = true
	case 256:
		fmt.Printf("%x\n", sha3.Sum256(data))
		success = true
	case 384:
		fmt.Printf("%x\n", sha3.Sum384(data))
		success = true
	case 512:
		fmt.Printf("%x\n", sha3.Sum512(data))
		success = true
	default:
		fmt.Printf("Unsupported hash length: %d\n", *hashlengthPtr)
	}
	if success && !*quiet {
		fmt.Printf("%s, %d bits\n", filename, *hashlengthPtr)
	}

	/*
		Arbitrary long hashes to be supported later

		// 64 bytes (512 bits) long hash provides 256-bit collision resistance.
		h := make([]byte, *hashlengthPtr)
		// Compute a 64-byte hash of buf and put it in h.
		sha3.ShakeSum256(h, data)
		fmt.Printf("%x\n", h)
	*/
}
