package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func main() {
	cmd := flag.String("dir", "", "")
	flag.Parse()
	
	dir := string(*cmd)
	
	println("Selected Dir is:", dir)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		rename(dir+f.Name(), dir+normalize(f.Name()))
	}
}

func normalize(fileName string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	
	newName, _, err := transform.String(t, fileName)
	
	if err != nil {
		log.Println(err)
	}
	// fmt.Println(string(newName))

	return string(newName)
}
// Mn is the set of Unicode characters in category Mn (Mark, nonspacing).
func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) || unicode.IsSpace(r)
}

func rename(oldName string, newName string) {
	err := os.Rename(oldName, newName)
	if err != nil {
		fmt.Println(err)
		return
	}
}
