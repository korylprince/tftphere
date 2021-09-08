package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	securejoin "github.com/cyphar/filepath-securejoin"
	"github.com/pin/tftp"
)

func main() {
	root := flag.String("root", ".", "root file directory for tftp server")
	force := flag.Bool("force", false, "allow files to be overwritten")
	flag.Parse()

	reader := func(fn string, rf io.ReaderFrom) error {
		fp, err := securejoin.SecureJoin(*root, fn)
		if err != nil {
			err = fmt.Errorf("could not find file path: %w", err)
			log.Printf("ERROR: reader: %s: %v\n", fn, err)
			return err
		}

		f, err := os.Open(fp)
		if err != nil {
			err = fmt.Errorf("could not open file: %w", err)
			log.Printf("ERROR: reader: %s: %v\n", fp, err)
			return err
		}

		n, err := rf.ReadFrom(f)
		if err != nil {
			err = fmt.Errorf("could not read file: %w", err)
			log.Printf("ERROR: reader: %s: %v\n", fp, err)
			return err
		}

		log.Printf("INFO: reader: %s: transfer complete (%d bytes)\n", fp, n)
		return nil
	}

	writer := func(fn string, wt io.WriterTo) error {
		fp, err := securejoin.SecureJoin(*root, fn)
		if err != nil {
			err = fmt.Errorf("could not find file path: %w", err)
			log.Printf("ERROR: writer: %s: %v\n", fn, err)
			return err
		}

		if info, err := os.Stat(fp); err == nil {
			if info.IsDir() {
				err = errors.New("directory exists with same name")
				log.Printf("ERROR: writer: %s: %v\n", fn, err)
				return err
			}

			if !(*force) {
				err = errors.New("file exists with same name")
				log.Printf("ERROR: writer: %s: %v\n", fn, err)
				return err
			}
		}

		f, err := os.Create(fp)
		if err != nil {
			err = fmt.Errorf("could not open file: %w", err)
			log.Printf("ERROR: writer: %s: %v\n", fp, err)
			return err
		}

		n, err := wt.WriteTo(f)
		if err != nil {
			err = fmt.Errorf("could not write file: %w", err)
			log.Printf("ERROR: writer: %s: %v\n", fp, err)
			return err
		}

		log.Printf("INFO: writer: %s: transfer complete (%d bytes)\n", fp, n)
		return nil
	}

	log.Printf("INFO: server: listening on :69 in %s\n", *root)
	s := tftp.NewServer(reader, writer)
	if err := s.ListenAndServe(":69"); err != nil {
		log.Printf("ERROR: server: unexpected shutdown: %v\n", err)
	}
}
