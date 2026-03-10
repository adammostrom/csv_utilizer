// Package main provides a command-line utility for appending data to CSV files.
// It supports reading data from text files and appending them as columns, rows, or tables
// to a target CSV file. Users can specify custom delimiters to parse input files.
//
// Usage:
//
//	csv_utilizer -target output.csv [-col file1.txt file2.txt] [-row rowdata.txt] [-table tabledata.txt] [-delimiter "delim"]
//
// Flags:
//
//	-col string
//	    Column files to append (can be repeated multiple times)
//	-row string
//	    File to write as a single row
//	-table string
//	    File to write as a table (multiple rows and columns)
//	-target string
//	    Target CSV file to append data to (required)
//	-delimiter string
//	    Delimiters to use when parsing input files (default: "\t\n;")
//
// The utility reads input files, parses them according to the specified format,
// and appends the parsed data to the target CSV file in append mode.
package main

/*
FLOW: Give it a csv as target, give it a text file with data
Choose if data is columns or rows
Append to the same csv


*/
import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type multiString []string

func (m *multiString) String() string {
	return fmt.Sprint("%v", *m)
}

// appends the value to the list, pointer so that its modified in place (same range).
// example use: -col file1.txt file2.txt file3.txt
func (m *multiString) Set(value string) error {
	*m = append(*m, value)
	return nil
}

func main() {

	var cols multiString
	flag.Var(&cols, "col", "column files (can repeat multiple times)")

	rowFile := flag.String("row", "", "Write as row")
	tableFile := flag.String("table", "", "Write as a table")
	target := flag.String("target", "", "target csv")
	delimFlag := flag.String("delimiter", "\t\n;", "delimiters")

	flag.Parse()

	delimiters := []rune(*delimFlag)

	if *target == "" {
		fmt.Println("missing -target")
		os.Exit(1)
	}

	for _, col := range cols {
		callCol(col, delimiters, *target)
	}

	if *rowFile != "" {
		callRow(*rowFile, delimiters, *target)
	}

	if *tableFile != "" {
		callTable(*tableFile, delimiters, *target)
	}

	fmt.Println("row:", *rowFile)
	fmt.Println("columns:", cols)
	fmt.Println("target:", *target)
	fmt.Println("delimiter:", *delimFlag)
}

// Reads from the file, parses data into table, writes table into csv.
func callTable(tableFile string, delimiters []rune, target string) {
	reads, err := readFromFile(tableFile)
	check(err)
	writeCSV(readTable(reads, delimiters), target)
}

// Reads from the file, parses data into columns, writes columns onto the csv (append mode).
func callCol(col string, delimiters []rune, target string) {
	reads, err := readFromFile(col)
	check(err)
	writeCSV(readColumn(reads, delimiters), target)
}

func callRow(rowFile string, delimiters []rune, target string) {
	reads, err := readFromFile(rowFile)
	check(err)
	data := readColumn(reads, delimiters)

	transposed, err := transpose(data)
	check(err)
	writeCSV(transposed, target)
}

func normalizeText(s string) string {
	s = strings.ReplaceAll(s, ",", ".")
	s = strings.ReplaceAll(s, "\n\r", "\n")
	s = strings.ReplaceAll(s, "\n\n", "\n")

	return s
}

func readFromFile(path string) (string, error) {

	data, err := os.ReadFile(path)
	check(err)

	if len(data) == 0 {
		return "", fmt.Errorf("file is empty")
	}

	s := normalizeText(string(data))

	return s, nil
}

func parseDelimiters(delim string) []rune {
	switch delim {
	case "tab", "\\t":
		return []rune{'\t'}
	case "newline", "\n":
		return []rune{'\n'}
	case "semicolon", "semicol", ";":
		return []rune{';'}
	case "none", "":
		return []rune{'\t', '\n', ';'} // Default to all of them

	default:
		return []rune(delim)
	}

}

// Read from file, then readColumn splits text into rows, each row is a single-column slice
// TODO: 2026-03-09: Implement delimiter for either many columns in the same file, or just general delimiters
func readColumn(s string, delimiters []rune) [][]string {

	dset := map[rune]bool{}

	// Adds all the delimiters given and sets them to true.
	for _, d := range delimiters {
		dset[d] = true
	}

	lines := strings.FieldsFunc(string(s), func(r rune) bool {
		return dset[r] // Return where dset[r] is true (the delimiters)
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

// Assume a delimiter separates the columns
// TODO: Add delimiter option
func readTable(s string, delimiters []rune) [][]string {

	dset := map[rune]bool{}

	for _, d := range delimiters {
		dset[d] = true
	}

	lines := strings.FieldsFunc(string(s), func(r rune) bool {
		return r == '\n' // Assume only newline separates rows
	})

	rows := [][]string{}
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l != "" {
			cols := strings.FieldsFunc(string(l), func(r rune) bool {
				return dset[r]
			})
			rows = append(rows, cols)
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

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Prepare to read the content of the csv file
	reader := csv.NewReader(file)

	// records now hold the read content
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	if len(records) != 0 {
		// Write data from the table to the records:
		appendCol(records, table)
	} else {
		records = append(records, table...)
	}

	out, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	writer := csv.NewWriter(out)
	writer.WriteAll(records)

	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
	writer.Flush()
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
