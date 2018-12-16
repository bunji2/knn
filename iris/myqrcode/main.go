package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	qrcode "github.com/skip2/go-qrcode"
)

const (
	usage = "%s text.txt\n"
)

func main() {
	os.Exit(run())
}

func run() int {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
		return 1
	}

	filePath := os.Args[1]

	outFilePath := filePath + ".QR.png"
	var bytes []byte
	var err error
	bytes, err = ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "run:", err)
		return 2
	}
	err = qrcode.WriteFile(string(bytes), qrcode.Medium, 256, outFilePath)

	if err != nil {
		fmt.Fprintln(os.Stderr, "run:", err)
		return 3
	}

	err = exec.Command("explorer", outFilePath).Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "run:", err)
		return 4
	}

	return 0
}
