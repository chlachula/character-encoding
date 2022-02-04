/*
MIT License

Copyright (c) 2022 Josef Chlachula

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NON INFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package character_encoding

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/htmlindex"
	"os"
)

func PrintEncodingNames() {
	fmt.Printf("List of canonical encoding names: ")
	comma := ""
	for _, enc := range charmap.All {
		name, _ := htmlindex.Name(enc)
		if len(name) > 2 {
			if name != "x-user-defined" {
				fmt.Printf("%s %s", comma, name)
			}
			comma = ","
		}
	}
	fmt.Printf(".\n")
}

func GetEncodingArgument(i int) (enc encoding.Encoding) {
	if len(os.Args) < i {
		HelpPrg("Missing encoding name")
	}
	enc, err := htmlindex.Get(os.Args[i])
	if err != nil {
		HelpPrg(err.Error())
	}
	return enc
}

// ReadLine returns a single line (without the ending \n) from the input buffered reader.
// An error is returned if there is an error with the buffered reader.
func ReadLine(r *bufio.Reader) (string, error) {
	var (
		isPrefix       = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func OpenFileArgument(i int) *os.File {
	if len(os.Args) < i {
		HelpPrg("Missing file name")
	}
	f, err := os.Open(os.Args[i])

	if err != nil {
		HelpPrg(fmt.Sprintf("error opening file: %v\n", err))

	}
	return f
}

func CreateFileArgument(i int) *os.File {
	if len(os.Args) < i {
		HelpPrg("Missing file name")
	}
	f, err := os.Create(os.Args[i])

	if err != nil {
		HelpPrg(fmt.Sprintf("error creating file: %v\n", err))
	}
	return f
}

func HelpPrg(errorMsg string) {
	exitCode := 0
	if errorMsg != "" {
		exitCode = 1
		fmt.Println(errorMsg)
	}
	//prg := os.Args[0]
	prg := "toUTF8"
	multiline := ` Character encoding conversion to UTF8 from another char encoding.
 (c) 2022-02-01 Josef Chlachula ver 0.1.0
 Utilize standard input output pipe and/or file input and output. 
Usage:
 %s -h ... HelpPrg text
 %s -a ... available char encodings
 %s -e ... encoding [default is Windows-1250]
 %s -@ ... utilizing stdin and stdout
 %s -f1 file1.txt -f2 file2.txt ... reading from file1.txt and writing to file2.txt
Examples:
 %s -e windows-1250 -@  # reads from stdin, writes to stdout
 %s -e windows-1250 -f1 in.txt f2 out-utf8.txt 
 %s -e windows-1250 -f1 in.txt #writes to stdout 
`
	fmt.Printf(multiline, prg, prg, prg, prg, prg, prg, prg, prg)
	os.Exit(exitCode)
}
