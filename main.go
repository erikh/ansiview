package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func errExit(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func main() {
	var f *os.File

	switch len(os.Args) {
	case 1:
		f = os.Stdin
	case 2:
		filename := os.Args[1]
		switch filename {
		case "-h", "-help", "--help":
			errExit("usage: ansiview <filename>")
		default:
		}

		var err error
		f, err = os.Open(filename)
		if err != nil {
			errExit(err.Error())
		}
		defer f.Close()
	default:
		errExit("invalid arguments")
	}

	dec := charmap.CodePage437.NewDecoder()

	for {
		buf := make([]byte, 4096)

		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			errExit(err.Error())
		}

		buf, err = dec.Bytes(buf[:n])
		if err != nil {
			errExit(err.Error())
		}

		os.Stdout.Write(buf)
	}
}
