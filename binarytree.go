package main

import "fmt"

type Node struct{
	Value int
	Left *Node
	Right *Node
}

func printNode(n *Node){
	fmt.Print("Value: ",n.Value)

	if n.Left != nil{
		fmt.Print(" Left: ",n.Left.Value)
	}

	if n.Right != nil{
		fmt.Print(" Right: ",n.Right.Value)
	}
	fmt.Println()

}

func main(){

	var N int
	fmt.Scanf("%d",&N)
	var nodes []Node = make([]Node,N)
	for i := 0; i < N; i++{
		var value, indexLeft, indexRight int
		fmt.Scanf("%d %d %d", &value, &indexLeft, &indexRight)
		//fmt.Println("Value: ", value, "Index Left: ", indexLeft, "Index Right: ", indexRight)
		nodes[i].Value = value
		if indexLeft >= 0{
			nodes[i].Left = &nodes[indexLeft]
		}
		if indexRight >= 0{
			nodes[i].Right = &nodes[indexRight]
		}

	}
	for _, node := range nodes{
		printNode(&node)
	}
}

