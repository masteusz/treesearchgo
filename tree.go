package main

import ("fmt")

type Node struct {
	Value int
	Cost int
	Children []Node
}

func (n Node) insert(val int, co int) *Node {
	n.Children.append(&Node{Value:val, Cost:co, Children: nil})
	return &n
}

func main() {
	fmt.Println("Works!")
}