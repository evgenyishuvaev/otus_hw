package main

import (
	"flag"
	"fmt"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func checkInputFlags() error {
	if from == "" || to == "" {
		return ErrMissingRequiredFlags
	}
	return nil
}

func main() {
	flag.Parse()
	err := checkInputFlags()
	if err != nil {
		err = fmt.Errorf("error: %w", err)
		fmt.Println(err)
		flag.Usage()
		return
	}
	Copy(from, to, offset, limit)
}
