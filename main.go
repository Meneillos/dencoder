package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func decode(str string) string {
	originalStr, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(-1)
	}
	return string(originalStr)
}

func usage() {
	fmt.Println("USAGE:")
	fmt.Println("dencoder [-encode / -decode] string")
	os.Exit(-1)
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	switch os.Args[1] {
	case "-encode":
		fmt.Println(encode(os.Args[2]))
	case "-decode":
		fmt.Println(decode(os.Args[2]))
	default:
		usage()
	}
}
