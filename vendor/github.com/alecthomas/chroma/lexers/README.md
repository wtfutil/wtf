# Lexer tests

The tests in this directory feed a known input `testdata/<name>.actual` into the parser for `<name>` and check
that its output matches `<name>.exported`.

## Running the tests

Run the tests as normal:
```go
go run ./lexers
```

## Updating the existing tests

You can regenerate all the test outputs

```go
RECORD=true go test ./lexers
```
