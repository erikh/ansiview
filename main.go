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
	if len(os.Args) != 2 {
		errExit("invalid arguments")
	}

	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		errExit(err.Error())
	}
	defer f.Close()

	buf := make([]byte, 4096)
	dec := charmap.CodePage437.NewDecoder()

	for {
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
