# CSV utilizer in go

Small tool for turning contents from text files into csv tables. 

This tool was created to simplify extracting data from Microsoft Word tables.
Copying Word tables into Excel often introduces formatting issues or data loss.
This utility allows copying table contents into text files and converting them
directly into CSV format, removing the Excel step entirely.

## Assumptions
The program assumes newline separates rows for the `-table` flag, the delimiter option will determine the data content for the columns

## How to run

Essentially copy the data from the entire table (or columns) into a text file. If no delimiter is selected, the default delimiter for tables will be TAB ('/t') for column separation and NEWLINE ('/n') for row separation.

Run the program:
`> go run main.go -target <file.csv> -table <data.txt>`

For adding column(s) to an existing csv file:

`> go run main.go -target <file.csv> -col <column.txt> -col <column2.txt> ...`

For adding a delimiter option (will apply to all file-arguments)

`> go run main.go -target <file.csv> -col <column.txt> -delimiter tab`

## Flags


Flags:
-col string
    Column files to append (can be repeated multiple times)
-row string
    File to write as a single row
-table string
    File to write as a table (multiple rows and columns)
-target string
    Target CSV file to append data to (required)
-delimiter string
    Delimiters to use when parsing input files (default: "\t\n;")

