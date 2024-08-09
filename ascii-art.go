package ascii

import (
	"fmt"
	"os"
	"strings"
)

func Ascii(fnt, str string) (string, bool ){
	Nstring := strings.ReplaceAll(str, "\r", "")      //delete carriage returns in str
	Nstring = strings.ReplaceAll(Nstring, "\n", `\n`) //replace new lines with literal new lines
	split := strings.Split(Nstring, `\n`)
	error500:= false
	text, err := os.ReadFile("ascii-art/" + fnt + ".txt")
	if err != nil {
		error500 = true 

		return "", error500
	}

	Nlines := strings.ReplaceAll(string(text), "\r", "") //delete carriage returns in font file
	lines := strings.Split(string(Nlines), "\n")
	output := ""
	for i := 0; i < len(split); i++ {
		if string(split[i]) == "" {
			fmt.Println()
		} else {
			for j := 0; j < 8; j++ {
				for k := 0; k < len(string(split[i])); k++ {
					output += lines[int(((rune(split[i][k])-32)*9+1))+j]
				}
				output += "\n"
			}
		}
	}
	return output, error500
}
