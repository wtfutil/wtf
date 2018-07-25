# Syntax Files

Here are highlights's syntax files.

Each yaml file specifies how to detect the filetype based on file extension or headers (first line of the file).
Then there are patterns and regions linked to highlight groups which tell micro how to highlight that filetype.

Making your own syntax files is very simple. I recommend you check the file after you are finished with the
[`syntax_checker.go`](./syntax_checker.go) program (located in this directory). Just place your yaml syntax
file in the current directory and run `go run syntax_checker.go` and it will check every file. If there are no
errors it will print `No issues!`.
