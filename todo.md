## TODO
[X] Add text file as argument input
[ ] Add text file parser into raw string
[X] Add transpose so its possible to turn into row
[ ] Add CSV preview stuff (amount of rows, fetcher etc)
	- Summary of numbers, summary of rows etc. Think Cloc but for csv
[/] Model and add the delimiter option  

[X] 2026-03-08: Fix so that the columns can be varying length, so each column that is appended just adds an empty string for the values it misses. 
- Fixed 2026-03-08. Currently column1 becomes a "base column".
	
[ ] - Add functionality to add values for a column (add column)
[ ] Add flags
	- [X] Add flag to specify column number like: -col 2 file1.txt
	- [X] Add flag to add row instead of col
	- [ ] Add delimiter option as flag:  --delimiter ","
[ ] Add proper error checks:
	- [ ] missing files
	- [ ] column mismatch
	- [ ] empty files
	- [ ] malformed CSV
[X] Add so that if csv doesnt exist, create one
[ ] Add functionality to accept stdin so you can call it like: `pbpbin | csvfill target.csv -col 2`  	
[ ] Add readme
[ ] Structure it properly and clean




## LOG

### 2026-03-08

Decide how the flow should be, should I call it with argument 1 be the source csv, the remaining arguments are the text files that contains the data, and a flag for row/col? Or should col/row flag be attachted to each file? Where should I add the headers? Should I add headers first or should they be automatically added? Should I assume the top row of the datafiles contain their respective header? How do I know to which column the data should be added to? Should I just assume iterative ordering?


Maybe for now, just have a CSV already made, assume headers exist, and just give it arguments as columns to add
