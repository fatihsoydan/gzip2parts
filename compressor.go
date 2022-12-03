package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

func closePart() {
	log.Printf("%s is ready\n", path.Join(OutputFolder, fmt.Sprintf("part.%04d", LastPartIndex)))
	CurrentPart.Close()
}

func createNewPart() {
	if CurrentPart != nil {
		closePart()
	}
	LastPosition = 0
	LastPartIndex++
	var err error
	newpartFileName := path.Join(OutputFolder, fmt.Sprintf("part.%04d", LastPartIndex))
	CurrentPart, err = os.Create(newpartFileName)
	if err != nil {
		log.Fatal(err)
	}
}

func addFile(fileInfo os.DirEntry, realpath string, archPath string) {
	var file FileDescriptor
	fi, _ := fileInfo.Info()
	file.ModifyDate = fi.ModTime()
	file.Name = fileInfo.Name()
	file.Folder = archPath
	buf := make([]byte, 0, 10*1024)
	f, err := os.Open(path.Join(realpath, fileInfo.Name()))
	if err != nil {
		log.Fatal(err)
	}
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	r := bufio.NewReader(f)
	for {
		n, err := r.Read(buf[:cap(buf)])
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		_, errw := w.Write(buf[:n])
		if errw != nil { // maybe memory problem
			log.Fatal(errw)
		}
	}
	w.Close()
	if CurrentPart == nil {
		createNewPart()
	}
	var part *FilePart
	var index = 0
	part = &FilePart{Index: index, Location: LastPartIndex}
	part.StartByte = LastPosition
	for {
		n, err := b.Read(buf[:cap(buf)])
		if LastPosition+int64(n) > Partsize {
			part.FinishByte = LastPosition
			if part.StartByte != part.FinishByte {
				file.Parts = append(file.Parts, *part)
				index++
			}
			createNewPart()
			part = &FilePart{Index: index, Location: LastPartIndex}
			part.StartByte = 0
		}
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				part.FinishByte = LastPosition
				file.Parts = append(file.Parts, *part)
				IndexObject = append(IndexObject, file)
				break
			}
			log.Fatal(err)
		}
		wb, _ := CurrentPart.Write(buf[:n])
		LastPosition += int64(wb)
	}
	TotalFiles++
}

func writeIndex() {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	enc.Encode(IndexObject)
	os.WriteFile(path.Join(OutputFolder, "index"), b.Bytes(), 0666)
}

func addFolder(fullpath string, archPath string) {
	files, err := os.ReadDir(fullpath)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() {
			addFolder(path.Join(fullpath, f.Name()), archPath+f.Name()+"/")
		} else if !strings.HasPrefix(f.Name(), ".") {
			addFile(f, fullpath, archPath)
		}
	}
}

func Compress() {
	starttime := time.Now()
	IndexObject = []FileDescriptor{}
	LastPosition, LastPartIndex, TotalFiles = 0, 0, 0

	addFolder(InputFolder, "/")

	closePart() // Close last part
	writeIndex()
	log.Printf("%d Files Compressed to %d parts in %v \n", TotalFiles, LastPartIndex, time.Since(starttime))
}
