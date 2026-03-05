package main

import (
	"fmt"
	"strings"
)

func main() {
	var test = `Internal Flash Checksum
	Internal SRAM Checksum 
	Internal SRAM Write Read 
	External SRAM Checksum 
	External SRAM Write Read 
	Co-Pro Red Communication`

	result := sanitizeString(test)
	fmt.Println(result)
}

// Backticks for strings spanning multiple lines

func readCol(col string) {

}

func sanitizeString(s string) string {

	//returnString := make([]rune, 0, len(s))

	// Iterate by RUNE, which is unicode safe (char iteration not safe for unicode)
	// i is the byte index
	//r is a rune (int32) containing the Unicode code point

    s = strings.ReplaceAll(s, ",", ".")
    s = strings.ReplaceAll(s, "\n", ",\n")

    // collapse multiple spaces into one
    //s = strings.Join(strings.Fields(s), " ")

    return s

}

/*


strings.ContainsRune(s, ',')
strings.ReplaceAll(s, ",", "")
strings.Split(s, ","


*/
