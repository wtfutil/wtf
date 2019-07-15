# ComputerUnits

[![GoDoc](https://godoc.org/github.com/hekmon/cunits?status.svg)](https://godoc.org/github.com/hekmon/cunits)

ComputerUnits allows to manipulate binary and decimal representations of bits and bytes.

## Constants example

```golang
fmt.Println(cunits.Kbit)
fmt.Println(cunits.Kibit)
fmt.Println(cunits.KB)
fmt.Println(cunits.KiB)
fmt.Printf("1000 MiB = %f MB\n", float64(1000)*cunits.MiB/cunits.MB)
```

will output:

```
1000
1024
8000
8192
1000 MiB = 1048.576000 MB
```

## Custom type example

```golang
size := cunits.Bit(58) * cunits.MiB
fmt.Println(size.Mbit())
fmt.Println(size.GiBString())
```

will output:

```
486.539264
0.06 GiB
```

## Parsing example

```golang
size, err := cunits.Parse("7632 MiB")
if err != nil {
    panic(err)
}
fmt.Println(size)
fmt.Println(size.KiB())
fmt.Println(size.KbitString())
```

will output:

```
7.45 GiB
7.815168e+06
64021856.26 Kbit
```
