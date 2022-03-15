package main

import (
	"os"
	"syscall"
)

type error interface {
	Error() string
}

// PathError records an error and the operation and
// file path that caused it.
type PathError struct {
	Op   string // "open", "unlink", etc
	Path string // The associated file.
	Err  error  // Returned by the system call.
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func readFile(fileName string) {
	for try := 0; try < 2; try++ {
		_, err := os.Create(fileName)
		if err == nil {
			return
		}
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
			// delete temp files
			continue
		}
		return
	}
}
