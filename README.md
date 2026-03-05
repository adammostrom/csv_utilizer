# CSV UTLILIZER IN GO

Idea: Take any text, decide the delimiter (tab, space, comma etc) from a text file, select which operation (to_column, to_row) and read the data from the textfile into a column or row, examples:


inside test.txt:

Some text that I want into a row

go run main.go test.txt -row

output :

some,text,that,I,want,into,a,row


go run main.go test.txt -col

some,
text,
that,
I,
want,
into,
a,
row,

The trick becomes to add columns to each other, so if I have one column with data like 1,2,3,4,5:

1,
2,
3,
4,
5,

and I want to add another column to this, like a,b,c,d,e, it should become:

1,a,
2,b,
3,c,
4,d,
5,e,


