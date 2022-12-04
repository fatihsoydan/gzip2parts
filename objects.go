package main

import "time"

type Index struct {
	Version       string
	PartCount     int
	ExtractedSize int64
	Files         []FileDescriptor
}

type FilePart struct {
	Index      int
	Location   int // PartFileIndex
	StartByte  int64
	FinishByte int64
}

type FileDescriptor struct {
	Name       string
	Folder     string
	ModifyDate time.Time
	Parts      []FilePart
}
