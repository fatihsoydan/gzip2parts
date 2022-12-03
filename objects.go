package main

import "time"

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
