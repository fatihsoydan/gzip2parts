package main

import (
	"flag"
	"log"
	"os"
)

var (
	Partsize      int64
	IndexObject   []FileDescriptor
	LastPartIndex int
	LastPosition  int64
	CurrentPart   *os.File
	InputFolder   string
	OutputFolder  string
	willCompress  bool
	TotalFiles    int
)

func init() {
	log.SetFlags(0)
	var compress bool
	var exract bool
	var psize int
	flag.BoolVar(&compress, "c", false, "Compress Folder")
	flag.BoolVar(&exract, "x", false, "Exract Folder")
	flag.StringVar(&InputFolder, "i", "", "InputFolder")
	flag.StringVar(&OutputFolder, "o", "", "OutputFolder")
	flag.IntVar(&psize, "ps", 1024*1000*3, "PartSize")
	flag.Parse()
	Partsize = int64(psize)
	if (!compress) && (!exract) {
		flag.Usage()
		log.Fatal("You must use -c or -e argument")
	}
	if (len(InputFolder) < 1) || (len(OutputFolder) < 1) {
		flag.Usage()
		log.Fatal("You must give InputFolder and OutputFolder")
	}
	willCompress = compress
}

func main() {
	if willCompress {
		Compress()
	} else {
		DeCompress()
	}
}
