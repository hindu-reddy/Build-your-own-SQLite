package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func main() {
	databaseFilePath := os.Args[1]
	command := os.Args[2]
	switch command {
	case ".dbinfo":
		databaseFile, err := os.Open(databaseFilePath)
		if err != nil {
			log.Fatal(err)
		}
		header := make([]byte, 100)
		_, err = databaseFile.Read(header)
		if err != nil {
			log.Fatal(err)
		}
		var pageSize uint16
		if err := binary.Read(bytes.NewReader(header[16:18]), binary.BigEndian, &pageSize); err != nil {
			fmt.Println("Failed to read integer:", err)
			return
		}
		// Print the page size
		fmt.Printf("database page size: %v", pageSize)
	default:
		fmt.Println("Unknown command", command)
		os.Exit(1)
	}
}
