package main

import (
	"encoding/csv"
	"os"
	"strings"
)

func main() {
	raw := `Internal Flash Checksum
Internal SRAM Checksum
Internal SRAM Write Read
External SRAM Checksum
Internal SRAM Write Read
Co-Pro Red Communication
asdasda
asd
adsdaasdad
asdasda
error,error
qwqwd
asdasdasd
`

	/* 	raw2 := `1
	   	2
	   	3
	   	4
	   	56
	   	` */
	raw3 := `A;B;C;D;E;F;G;H;I;J;K;L;M;N;O;P;Q;R;S;T;U;V;W;X;Y;Z;Å;Ä;Ö`

	// Step 1: read as column
	//newspaceDelimiters := []string{"\r\n", "\n"}
	//semiColonDelimiters := []string{";", ";"}

	col := readColumn(normalizeText(raw))
	//col1 := readColumn(raw2)
	col2 := readColumn(normalizeText(raw3))

	// Step 2: append a duplicate column (for demo)
	table := appendCol(col, col2, col)

	// Step 3: write CSV
	writeCSV(table, "output.csv")
}

func normalizeText(s string) string {
	s = strings.ReplaceAll(s, ",", ".")
	s = strings.ReplaceAll(s, "\n\r", "\n")
	s = strings.ReplaceAll(s, "\n\n", "\n")

	return s
}

// readColumn splits text into rows, each row is a single-column slice
func readColumn(s string) [][]string {
	//lines := strings.Split(strings.ReplaceAll(s, delimiter[0], delimiter[1]), delimiter[1])
	lines := strings.FieldsFunc(s, func(r rune) bool {
		return r == '\n' || r == '\t' || r == ';'
	})
	rows := [][]string{}
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l != "" {
			rows = append(rows, []string{l})
		}
	}
	return rows
}

// appendCol appends additional columns to each row
func appendCol(trg [][]string, srcs ...[][]string) [][]string {
	rows := len(trg)
	if rows == 0 && len(srcs) > 0 {
		rows = len(srcs[0])
	}

	// ensure trg has enough rows
	for len(trg) < rows {
		trg = append(trg, []string{})
	}

	for _, col := range srcs {
		for i := 0; i < rows; i++ {
			val := ""
			if i < len(col) {
				val = col[i][0]
			}
			trg[i] = append(trg[i], val)
		}
	}

	return trg
}

// writeCSV writes [][]string to a CSV file
func writeCSV(table [][]string, path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range table {
		writer.Write(row)
	}
}
