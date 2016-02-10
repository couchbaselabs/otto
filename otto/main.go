package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/couchbaselabs/otto"
	"github.com/couchbaselabs/otto/underscore"
)

var underscoreFlag *bool = flag.Bool("underscore", true, "Load underscore into the runtime environment")

func main() {
	flag.Parse()
	var script []byte
	var err error
	filename := flag.Arg(0)
	if filename == "" || filename == "-" {
		script, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Printf("Can't read stdin: %v\n", err)
			os.Exit(64)
		}
	} else {
		script, err = ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("Can't open file \"%v\": %v\n", filename, err)
			os.Exit(64)
		}
	}
	if !*underscoreFlag {
		underscore.Disable()
	}
	Otto := otto.New()
	_, err = Otto.Run(string(script))
	if err != nil {
		fmt.Println(err)
		os.Exit(64)
	}
}
