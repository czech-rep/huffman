package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
)

type UnhandledCharacter string

func (char UnhandledCharacter) Error() string {
	return fmt.Sprintf("Unhandled character: %v", string(char))
}

type HuffmanError string

func (msg HuffmanError) Error() string {
	return string(msg)
}

type Node struct {
	Value  string `json:"value"` // nonempty string indicates a leaf
	Weight int    `json:"weight"`
	Lchild *Node  `json:"lchild,omitempty"`
	Rchild *Node  `json:"rchild,omitempty"`
}

func (root *Node) Json() (string, error) {
	jsonData, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func Decode(root *Node, message string) (string, error) {
	placeholder := []string{}
	result, err := decode(stringToSlice(message), root, root, placeholder)
	if err != nil {
		return "", err
	}
	return sliceToString(result), nil
}

func decode(message []string, root *Node, node *Node, result []string) ([]string, error) {
	// code contains "1" and "0"
	if len(message) == 0 {
		if node != root {
			return result, HuffmanError("Finished decoding inside tree")
		}
		return result, nil
	}
	bit, message := message[0], message[1:]

	var next_node *Node

	if bit == "1" {
		next_node = node.Lchild
	} else if bit == "0" {
		next_node = node.Rchild
	} else {
		return result, UnhandledCharacter(bit)
	}
	if next_node.Value != "" {
		// reached leaf. append letter and go back to root
		return decode(message, root, root, append(result, next_node.Value))
	} else {
		return decode(message, root, next_node, result)
	}
}
func BuildTree(weightedStrings map[string]int) *Node { // TODO func to update tree
	nodeHeap := &NodeHeap{}
	heap.Init(nodeHeap)

	for word, weight := range weightedStrings {
		heap.Push(nodeHeap, Node{word, weight, nil, nil})
	}

	for len(*nodeHeap) > 1 {
		min1 := heap.Pop(nodeHeap).(Node)
		min2 := heap.Pop(nodeHeap).(Node)
		heap.Push(
			nodeHeap,
			Node{"", min1.Weight + min2.Weight, &min1, &min2},
		)
	}
	result := heap.Pop(nodeHeap).(Node)
	return &result
}
func getHuffmanMap(root *Node) map[string]string {
	// mapping from symbols to Huffman codes
	result := make(map[string]string)
	nodes := []*Node{root}
	paths := []string{""}
	var elem *Node
	var path string

	for len(nodes) > 0 { // depth first search
		length := len(nodes)
		nodes, elem = nodes[:length-1], nodes[length-1]
		paths, path = paths[:length-1], paths[length-1]
		if elem.Value != "" {
			result[elem.Value] = path
		} else {
			nodes = append(nodes, elem.Lchild)
			paths = append(paths, path+"1")

			nodes = append(nodes, elem.Rchild)
			paths = append(paths, path+"0")
		}
	}
	return result
}
func Encode(root *Node, input string) (string, error) {
	huffmanMap := getHuffmanMap(root)
	var result string = ""
	for _, item := range stringToSlice(input) {
		if code, ok := huffmanMap[item]; ok {
			result = result + code
		} else {
			return "", UnhandledCharacter(item)
		}
	}
	return result, nil
}
