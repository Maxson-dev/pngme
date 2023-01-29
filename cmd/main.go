package main

import (
	"flag"
	"fmt"
	"os"
	"pngme/internal/api"
)

const (
	encode = "encode"
	decode = "decode"
	print  = "print"
	remove = "remove"
)

func main() {
	encodeCmd := flag.NewFlagSet(encode, flag.ExitOnError)
	encFile := encodeCmd.String("file", "", "png file to encode message")
	encHead := encodeCmd.String("head", "", "message header")
	encMsg := encodeCmd.String("msg", "", "message to encoding")

	decodeCmd := flag.NewFlagSet(decode, flag.ExitOnError)
	decFile := decodeCmd.String("file", "", "png file for decode message")
	decHead := decodeCmd.String("head", "", "message header")

	printCmd := flag.NewFlagSet(print, flag.ExitOnError)
	printFile := printCmd.String("file", "", "png file to print")

	rmCmd := flag.NewFlagSet(remove, flag.ExitOnError)
	rmFile := rmCmd.String("file", "", "png file for remove message")
	rmHead := rmCmd.String("head", "", "message header")

	switch os.Args[1] {
	case encode:
		encodeCmd.Parse(os.Args[2:])
		err := api.EncodeMessage(*encFile, *encHead, *encMsg)
		if err != nil {
			fmt.Printf("Could not encode message: %v\n", err)
			os.Exit(1)
		}
	case decode:
		decodeCmd.Parse(os.Args[2:])
		err := api.DecodeMessage(*decFile, *decHead)
		if err != nil {
			fmt.Printf("Could not decode message: %v\n", err)
			os.Exit(1)
		}
	case print:
		printCmd.Parse(os.Args[2:])
		err := api.PrintPngFile(*printFile)
		if err != nil {
			fmt.Printf("Could not print file %v\n", err)
			os.Exit(1)
		}
	case remove:
		rmCmd.Parse(os.Args[2:])
		err := api.RemoveMessage(*rmFile, *rmHead)
		if err != nil {
			fmt.Printf("Could not remove message %s: %v\n", *rmHead, err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command %s\n", os.Args[1])
		os.Exit(1)
	}

}
