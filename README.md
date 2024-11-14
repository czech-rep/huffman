Huffman coding in go. usage:

## install dependencies
```
go mod download
```

## create executable
```
go build -o huffman
```
(to use instead of `go run .` or install in PATH)

## encode a file
```
go run . encode <file>
```
will produce binary file `<file>.huffman` and json representation of tree `<file>.tree.huffman`

## decode
```
go run . --tree <file>.tree.huffman decode <file>.huffman
```

## run tests
```
go test
```

huffman tree example:
<pre>
        x
    x       x
 a    b  c     d

11   10  01    00
</pre>

encoding: abac -> 11101101
done with a map: {a: 11, b:10, ...}

decoding 000111 -> 00 01 11 -> d c a
done by recursiwely traversing graph

desired workflow - encoding a file
provided file path, you parse by characters
-> map[string]int32 - words with weights -> huffman tree -> huffman encoding -> save to binary file
