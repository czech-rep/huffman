package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func DecodeCmd(inputFile, treeFile string) (string, error) {
	root, err := ReadTreeFromFile(treeFile)
	if err != nil {
		return "", err
	}
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return "", err
	}
	bitString, err := BytesToBits(content)
	if err != nil {
		return "", err
	}
	result, err := Decode(root, bitString)
	if err != nil {
		return "", err
	}
	return result, nil
}
func EncodeCmd(inputFile, outFile, outTreeFile string) error {
	content, err := ReadTextFile(inputFile)
	if err != nil {
		return err
	}
	var root *Node
	contedStrings := Counted(stringToSlice(content))
	root = BuildTree(contedStrings)

	result, err := Encode(root, content)
	if err != nil {
		return err
	}
	resultBytes := BinaryToBytes(result)
	WriteTreeToFile(outTreeFile, root)
	os.WriteFile(outFile, resultBytes, 0644)
	return nil
}

func EncodeWithTree(inputFile, treeFile, outFile string) error {
	content, err := ReadTextFile(inputFile)
	if err != nil {
		return err
	}
	var root *Node
	root, err = ReadTreeFromFile(treeFile)
	if err != nil {
		return err
	}
	result, err := Encode(root, content)
	if err != nil {
		return err
	}
	WriteToFile(outFile, result)
	return nil
}

func main() {
	var mode, inputFile, treeFile string

	app := &cli.App{
		Name:  "Huffman coding",
		Usage: "[-tree <tree_path>] encode <path>\n-tree <tree_path> decode <path>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "tree",
				Value:       "",
				Usage:       "path to tree file",
				Destination: &treeFile,
			},
		},
		Action: func(ctx *cli.Context) error {
			if ctx.Args().Len() < 2 {
				return cli.Exit("Too few arguments", 1)
			}
			mode = ctx.Args().Get(0)
			inputFile = ctx.Args().Get(1)

			if mode == "decode" {
				if treeFile == "" {
					return cli.Exit("Error: --tree is required for decoding", 1)
				}
				result, err := DecodeCmd(inputFile, treeFile)
				if err != nil {
					return err
				}
				fmt.Println(result)
			} else if mode == "encode" {
				outFile := inputFile + ".huffman"
				outTreeFile := inputFile + ".tree.huffman"
				if treeFile == "" {
					err := EncodeCmd(inputFile, outFile, outTreeFile)
					if err != nil {
						return err
					}
				} else {
					err := EncodeWithTree(inputFile, treeFile, outFile)
					if err != nil {
						return err
					}
				}
				fmt.Printf("Encoded to %v\n", outFile)
			} else {
				return cli.Exit("mode may be encode / decode", 1)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
