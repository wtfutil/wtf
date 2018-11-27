bytefmt
=======

**Note**: This repository should be imported as `code.cloudfoundry.org/bytefmt`.

Human-readable byte formatter.

Example:

```go
bytefmt.ByteSize(100.5*bytefmt.MEGABYTE) // returns "100.5M"
bytefmt.ByteSize(uint64(1024)) // returns "1K"
```

For documentation, please see http://godoc.org/code.cloudfoundry.org/bytefmt
