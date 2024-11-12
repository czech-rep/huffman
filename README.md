Huffman coding in go. usage:

## encode a file
```
go run . encode <file>
```
will produce binary file `<file>.huffman` and json representation of tree `<file>.tree.huffman`

## decode
```
go run . --tree <file>.tree.huffman decode <file>.huffman
```

how works:

        x
    x       x
 a    b  c     d

11   10  01    00

encoding: abac -> 11101101
will use a mapping: {a: 11, ...}

decoding 000111 -> dca
done by recursiwely traversing graph

desired workflow - encoding a file
provided file path, you parse by characters
-> map[string]int32 - words with weights -> huffman tree -> huffman encoding -> save to binary file
