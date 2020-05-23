package main

import (
	"flag"
	"io"
	"log"
	"os"
	"sync"
)

var (
	dstFilename string
)

func init() {
	flag.StringVar(&dstFilename, "o", "-", "output filepath(default is stdout)")
}

func main() {
	flag.Parse()
	args := flag.Args()
	filenames := args

	var dstWriter io.Writer
	if dstFilename == "-" {
		dstWriter = os.Stdout
	} else {
		f, err := os.Create(dstFilename)
		if err != nil {
			log.Fatalf("[Error]: create file '%s': %s", dstFilename, err)
		}
		defer f.Close()
		dstWriter = f
	}

	wg := &sync.WaitGroup{}
	for _, filename := range filenames {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			var srcReader io.Reader
			if filename != "-" {
				f, err := os.Open(filename)
				if err != nil {
					log.Fatalf("[Error]: open file '%s': %s", filename, err)
				}
				defer f.Close()
				srcReader = f
			} else {
				srcReader = os.Stdin
			}
			io.Copy(dstWriter, srcReader)
		}(filename)
	}
	wg.Wait()
}
