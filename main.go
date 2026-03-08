package main

/*
FLOW: Give it a csv as target, give it a text file with data
Choose if data is columns or rows
Append to the same csv


*/
import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file.txt> <file.txt> ... ")
		os.Exit(1)
	}

	for _, arg := range os.Args[1:] {
		writeCSV((readColumn(readFromFile(arg))), "output.csv")
	}

	// Step 1: read as column
	//newspaceDelimiters := []string{"\r\n", "\n"}
	//semiColonDelimiters := []string{";", ";"}

	/* 	col1 := readFromFile("text.txt")
	   	col2 := readFromFile("col2.txt")
	   	col3 := readFromFile("col3.txt")

	   	col_1 := readColumn(col1)
	   	col_2 := readColumn(col2)
	   	col_3 := readColumn(col3)

	   	table := appendCol(col_1, col_2, col_3)

	   	writeCSV(table, "output.csv") */

}

func normalizeText(s string) string {
	s = strings.ReplaceAll(s, ",", ".")
	s = strings.ReplaceAll(s, "\n\r", "\n")
	s = strings.ReplaceAll(s, "\n\n", "\n")

	return s
}

func readFromFile(path string) string {

	data, err := os.ReadFile(path)
	check(err)

	s := normalizeText(string(data))

	return s
}

// Read from file, then readColumn splits text into rows, each row is a single-column slice
func readColumn(s string) [][]string {

	//lines := strings.Split(strings.ReplaceAll(s, delimiter[0], delimiter[1]), delimiter[1])
	lines := strings.FieldsFunc(string(s), func(r rune) bool {
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
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	// Prepare to read the content of the csv file
	reader := csv.NewReader(file)

	// records now hold the read content
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	if len(records) != 0 {
		// Write data from the table to the records:
		appendCol(records, table)
	} else {
		records = append(records, table...)
	}

	out, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	writer := csv.NewWriter(out)
	writer.WriteAll(records)

	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
}

// Turn cols into rows, and rows into cols
func transpose(matrix [][]string) ([][]string, error) {

	if len(matrix) == 0 {
		return nil, nil
	}

	for i, _ := range matrix {
		if len(matrix[i]) != len(matrix[0]) {
			return nil, fmt.Errorf("Matrix is not rectangular, irregluar length for: %v", matrix[i])
		}
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
