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

var encodingNames = []string{
	"CodePage037",
	"CodePage1047",
	"CodePage1140",
	"CodePage437",
	"CodePage850",
	"CodePage852",
	"CodePage855",
	"CodePage858",
	"CodePage860",
	"CodePage862",
	"CodePage863",
	"CodePage865",
	"CodePage866",
	"ISO8859_1",
	"ISO8859_10",
	"ISO8859_13",
	"ISO8859_14",
	"ISO8859_15",
	"ISO8859_16",
	"ISO8859_2",
	"ISO8859_3",
	"ISO8859_4",
	"ISO8859_5",
	"ISO8859_6",
	"ISO8859_7",
	"ISO8859_8",
	"ISO8859_9",
	"KOI8R",
	"KOI8U",
	"Macintosh",
	"MacintoshCyrillic",
	"Windows1250",
	"Windows1251",
	"Windows1252",
	"Windows1253",
	"Windows1254",
	"Windows1255",
	"Windows1256",
	"Windows1257",
	"Windows1258",
	"Windows874",
}

func printEncodingNamesOld() {
	for i, e := range encodingNames {
		fmt.Printf("%d:%s, ", i, e)
	}
	fmt.Printf("\n")
}

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
func getDecoderOld(name string) *encoding.Decoder {

	var d *encoding.Decoder
	switch {
	case name == "CodePage037":
		d = charmap.CodePage037.NewDecoder()
	case name == "CodePage1047":
		d = charmap.CodePage1047.NewDecoder()
	case name == "CodePage1140":
		d = charmap.CodePage1140.NewDecoder()
	case name == "CodePage437":
		d = charmap.CodePage437.NewDecoder()
	case name == "CodePage850":
		d = charmap.CodePage850.NewDecoder()
	case name == "CodePage852":
		d = charmap.CodePage852.NewDecoder()
	case name == "CodePage855":
		d = charmap.CodePage855.NewDecoder()
	case name == "CodePage858":
		d = charmap.CodePage858.NewDecoder()
	case name == "CodePage860":
		d = charmap.CodePage860.NewDecoder()
	case name == "CodePage862":
		d = charmap.CodePage862.NewDecoder()
	case name == "CodePage863":
		d = charmap.CodePage863.NewDecoder()
	case name == "CodePage865":
		d = charmap.CodePage865.NewDecoder()
	case name == "CodePage866":
		d = charmap.CodePage866.NewDecoder()
	case name == "ISO8859_1":
		d = charmap.ISO8859_1.NewDecoder()
	case name == "ISO8859_10":
		d = charmap.ISO8859_10.NewDecoder()
	case name == "ISO8859_13":
		d = charmap.ISO8859_13.NewDecoder()
	case name == "ISO8859_14":
		d = charmap.ISO8859_14.NewDecoder()
	case name == "ISO8859_15":
		d = charmap.ISO8859_15.NewDecoder()
	case name == "ISO8859_16":
		d = charmap.ISO8859_16.NewDecoder()
	case name == "ISO8859_2":
		d = charmap.ISO8859_2.NewDecoder()
	case name == "ISO8859_3":
		d = charmap.ISO8859_3.NewDecoder()
	case name == "ISO8859_4":
		d = charmap.ISO8859_4.NewDecoder()
	case name == "ISO8859_5":
		d = charmap.ISO8859_5.NewDecoder()
	case name == "ISO8859_6":
		d = charmap.ISO8859_6.NewDecoder()
	case name == "ISO8859_7":
		d = charmap.ISO8859_7.NewDecoder()
	case name == "ISO8859_8":
		d = charmap.ISO8859_8.NewDecoder()
	case name == "ISO8859_9":
		d = charmap.ISO8859_9.NewDecoder()
	case name == "KOI8R":
		d = charmap.KOI8R.NewDecoder()
	case name == "KOI8U":
		d = charmap.KOI8U.NewDecoder()
	case name == "Macintosh":
		d = charmap.Macintosh.NewDecoder()
	case name == "MacintoshCyrillic":
		d = charmap.MacintoshCyrillic.NewDecoder()
	case name == "Windows1250":
		d = charmap.Windows1250.NewDecoder()
	case name == "Windows1251":
		d = charmap.Windows1251.NewDecoder()
	case name == "Windows1252":
		d = charmap.Windows1252.NewDecoder()
	case name == "Windows1253":
		d = charmap.Windows1253.NewDecoder()
	case name == "Windows1254":
		d = charmap.Windows1254.NewDecoder()
	case name == "Windows1255":
		d = charmap.Windows1255.NewDecoder()
	case name == "Windows1256":
		d = charmap.Windows1256.NewDecoder()
	case name == "Windows1257":
		d = charmap.Windows1257.NewDecoder()
	case name == "Windows1258":
		d = charmap.Windows1258.NewDecoder()
	case name == "Windows874":
		d = charmap.Windows874.NewDecoder()
	default:
		HelpPrg("Unknown encoding")
	}
	return d
}

// Readline  returns a single line (without the ending \n) from the input buffered reader.
// An error is returned if there is an error with the buffered reader.
func ReadLine(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
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
