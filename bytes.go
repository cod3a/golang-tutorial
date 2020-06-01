package main

import (
	"bytes"
	"fmt"
)

func lowerCaseInBytes() {
	myByte := []byte("Eisson Alipio")
	var byteString = bytes.ToLower(myByte) // [101 105 115 115 111 110 32 97 108 105 112 105 111]
	fmt.Printf("%s\n", byteString)

}
