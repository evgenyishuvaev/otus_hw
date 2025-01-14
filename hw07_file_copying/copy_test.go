package main

import "testing"

func TestCopy(t *testing.T) {
	// Place your code here.
	//
	//
	//
	t.Run("missing required flags", func(t *testing.T) {})
	t.Run("offset is less then 0", func(t *testing.T) {})
	t.Run("offset is bigger then file size", func(t *testing.T) {})

	t.Run("limit is less then 0", func(t *testing.T) {})

	t.Run("copy file with unknown filesize(like /dev/urandom)", func(t *testing.T) {})

	t.Run("src and dst files is equal", func(t *testing.T) {})
	t.Run("wrong permissions", func(t *testing.T) {})
	t.Run("destination file has extra bytes at end of file by chuncsize")
}
