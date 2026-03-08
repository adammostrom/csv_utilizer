## TODO
[ ] Add text file as argument input
[ ] Add text file parser into raw string
[ ] Add transpose so its possible to turn into row
[ ] Add flags for turning contents of textfile into col or row
[ ] Add CSV preview stuff (amount of rows, fetcher etc)
	- Summary of numbers, summary of rows etc. Think Cloc but for csv
[/] Model and add the delimiter option  

[X] - 2026-03-08: Fix so that the columns can be varying length, so each column that is appended just adds an empty string for the values it misses. 
- Fixed 2026-03-08. Currently column1 becomes a "base column".
	




## LOG

### 2026-03-08

Decide how the flow should be, should I call it with argument 1 be the source csv, the remaining arguments are the text files that contains the data, and a flag for row/col? Or should col/row flag be attachted to each file? Where should I add the headers? Should I add headers first or should they be automatically added? Should I assume the top row of the datafiles contain their respective header? How do I know to which column the data should be added to? Should I just assume iterative ordering?


Maybe for now, just have a CSV already made, assume headers exist, and just give it arguments as columns to add