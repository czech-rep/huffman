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
func EncodeCmd(inputFile, treeFile string) (string, error) {
	content, err := ReadTextFile(inputFile)
	if err != nil {
		return "", err
	}
	var root *Node
	if treeFile != "" {
		root, err = ReadTreeFromFile(treeFile)
		if err != nil {
			return "", err
		}
	} else {
		contedStrings := Counted(stringToSlice(content))
		root = BuildTree(contedStrings)
		WriteTreeToFile(inputFile+".tree.huffmann", root) // somewhere else
	}
	result, err := Encode(root, content)
	if err != nil {
		return "", err
	}
	return result, nil
}

func main() {
	var mode, inputFile, treeFile string

	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
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
				fmt.Println("invalid args")
			}
			mode = ctx.Args().Get(0)
			inputFile = ctx.Args().Get(1)

			if mode == "decode" {
				if treeFile == "" {
					log.Fatalf("Error: --tree is required for recoding")
				}
				result, err := DecodeCmd(inputFile, treeFile)
				if err != nil {
					return err
				}
				fmt.Println(result)
			} else if mode == "encode" {
				result, err := EncodeCmd(inputFile, treeFile)
				if err != nil {
					return err
				}
				outPath := inputFile + ".huffman"
				os.WriteFile(outPath, BinaryToBytes(result), 0644)
				fmt.Printf("Encoded to %v\n", outPath)
			} else {
				return cli.Exit("invalid mode", 1)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
