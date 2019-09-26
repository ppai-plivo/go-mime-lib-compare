package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"text/tabwriter"

	"github.com/gabriel-vasile/mimetype"
	"github.com/h2non/filetype"
	"github.com/rakyll/magicmime"
	"github.com/vimeo/go-magic/magic"
	"github.com/zRedShift/mimemagic"
)

var (
	tabW = tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
)

func printContentTypes(testDir string) error {

	if err := magicmime.Open(magicmime.MAGIC_MIME_TYPE); err != nil {
		return err
	}

	defer tabW.Flush()

	fmt.Fprintf(tabW, "## MIME detection capablity comparison\n\n")

	fmt.Fprintf(tabW, "|%s\t%s\t%s\t%s\t%s\t%s\t%s\t\n",
		"File path", "http.DetectContentType", "gabriel-vasile/mimetype", "zRedShift/mimemagic", "rakyll/magicmime", "vimeo/go-magic", "h2non/filetype")
	fmt.Fprintf(tabW, "|---\t---\t---\t---\t---\t---\t---\t\n")

	return filepath.Walk(testDir, detectContentType)
}

func detectContentType(path string, info os.FileInfo, err error) error {

	if err != nil {
		fmt.Printf("Error accessing path %s: %v\n", path, err)
		return err
	}

	if info.IsDir() {
		return nil
	}

	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", path, err)
		return err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("ioutil.ReadAll() failed for path %s: %v\n", path, err)
		return err
	}

	ctype1 := http.DetectContentType(b)
	ctype2, _ := mimetype.Detect(b)
	ctype3 := mimemagic.MatchMagic(b).MediaType()
	ctype4, _ := magicmime.TypeByBuffer(b)
	ctype5 := magic.MimeFromBytes(b)
	mtype, _ := filetype.Match(b)
	ctype6 := mtype.MIME.Type + "/" + mtype.MIME.Subtype

	fmt.Fprintf(tabW, "|%s\t%s\t%s\t%s\t%s\t%s\t%s\t\n", path, ctype1, ctype2, ctype3, ctype4, ctype5, ctype6)

	return nil
}
