package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrMissingRequiredFlags  = errors.New("missing required flags")
	chunkSize                = 64
)

func copyAll(srcFile *os.File, dstFile *os.File) error {
	buf := make([]byte, chunkSize)
	for {
		readedChunk, readError := srcFile.Read(buf)
		if readError != nil && readError != io.EOF {
			fmt.Println(readError)
			return readError
		}

		if errors.Is(readError, io.EOF) {
			return nil
		}

		if len(buf) > readedChunk {
			buf = buf[:readedChunk+1]
		}

		_, writeError := dstFile.Write(buf)
		if writeError != nil {
			return writeError
		}
	}
}

func copyPartial(srcFile *os.File, dstFile *os.File, offset, limit int64) error {
	buf := make([]byte, chunkSize)
	for {
		readedN, err := os.ReadFile()

	}
	return nil
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	var copyError error

	srcFile, err := os.Open(fromPath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	srcFileInfo, err := srcFile.Stat()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if srcFileInfo.Size() < offset {
		fmt.Println(ErrOffsetExceedsFileSize)
		return ErrOffsetExceedsFileSize
	}
	defer srcFile.Close()

	dstFile, err := os.Create(toPath)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer dstFile.Close()

	if offset == 0 && limit == 0 {
		copyError = copyAll(srcFile, dstFile)
		if copyError != nil {
			fmt.Println(copyError)
			return copyError
		}
	} else {
		copyError = copyPartial(srcFile, dstFile, offset, limit)
		if copyError != nil {
			fmt.Println(copyError)
			return copyError
		}
	}
	return nil
}
