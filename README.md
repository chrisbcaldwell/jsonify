# Jsonify

This repository contains a command line tool for converting a .csv file into a JSON lines file.

## Clone the project

```
$ git clone https://github.com/chrisbcaldwell/jsonify
$ cd jsonify
```

The repository is located at https://github.com/chrisbcaldwell/jsonify.

## jsonify

```
$ go build
$ jsonify -path
```

A command line tool that accepts a file path after the `-path` flag.  The file path can be relative to the directory in which jsonify is saved.

The output is a JSON lines file.  For more on JSON lines, see the [JSON LInes documentation](https://jsonlines.org/).  The resulting file is saved to the same directory as the original .csv file with the extension .jsonl appended to the file name.

Example usage with [test1.csv](https://github.com/chrisbcaldwell/jsonify/blob/main/testdata/test1.csv):

```
$ jsonify -path ./testdata/test1.csv
JSON file saved at ./testdata/test1.csv.jsonl
```

### Technical Details

* Comma delimiters are supported.
* A header row must be used.
* The number of fields in each row must match the number of fileds in the header.
* JSON lines output fields are formatted as strings.  Further processing is needed to convert numeric fields.
* Row order from the original file is retained.  Field order is not preserved in the .jsonl output.

### Future Update Possibilities

In no particular order, future updates could include:
* Support for alternative delimiters in the .csv file.
* Processing without a header row.
* Retaining numeric data types.
* Preservation of field order from the original .csv file.

### References

Andile, Maximilien n.d.  "How to read a file line by line with Go."  Accessed April 16, 2025.  https://www.practical-go-lessons.com/post/how-to-read-a-file-line-by-line-with-go-cbmaj47rttts70kq7c9g

Gerardi, Ricardo 2021.  *Powerful Command-Line Applications in Go,* Chapter 3.  Pragmatic Bookshelf

Ramanathan, Naveen 2023.  "Writing Files using Go."  *golangbot.com* (blog), August 13, 2023.  https://golangbot.com/write-files/

SyntaxRules 2019.  "Reading CSV file in Go."  stackoverflow, November 13, 2019.  https://stackoverflow.com/a/58841827


