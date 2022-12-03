package main

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"time"
)

var maxPartsUsed int

func extractFile(file FileDescriptor) {
	var gzipContent bytes.Buffer
	sort.Slice(file.Parts, func(i, j int) bool {
		return file.Parts[i].Index < file.Parts[j].Index
	})

	for _, fp := range file.Parts {
		if fp.Location > maxPartsUsed {
			maxPartsUsed = fp.Location
		}
		fsource, err := os.Open(path.Join(InputFolder, fmt.Sprintf("part.%04d", fp.Location)))
		if err != nil {
			log.Fatal(err)
		}
		fsource.Seek(int64(fp.StartByte), 0)
		var b []byte = make([]byte, fp.FinishByte-fp.StartByte)
		fsource.Read(b)
		gzipContent.Write(b)
		fsource.Close()
	}
	buf := bytes.NewBuffer(gzipContent.Bytes())
	r, err := gzip.NewReader(buf)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	fileContents, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	folder2write := path.Join(OutputFolder, file.Folder)
	if _, err := os.Stat(folder2write); os.IsNotExist(err) {
		os.MkdirAll(folder2write, os.ModePerm)
	}
	ioutil.WriteFile(path.Join(folder2write, file.Name), fileContents, 0666)
	TotalFiles++
}

func readIndex() {
	b, _ := os.ReadFile(path.Join(InputFolder, "index"))
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	dec.Decode(&IndexObject)
}

func DeCompress() {
	starttime := time.Now()
	maxPartsUsed = 0
	readIndex()
	for _, obj := range IndexObject {
		extractFile(obj)
	}

	log.Printf("%d Files Extracted from %d parts in %v \n", TotalFiles, maxPartsUsed, time.Since(starttime))
}
