package main

import "fmt"

type Node struct{
	Key int
	Left *Node
	Right *Node
}

// Insert will add a node to the tree
// The key to add should not be already in the tree
func(n *Node) Insert(k int){
	if(k>n.Key){
		// Move right
		if n.Right==nil{
			n.Right=&Node{Key:k}
		}else{
			n.Right.Insert(k)
		}
	}else{
		// Mov left
		if n.Left==nil{
			n.Left=&Node{Key:k}
		}else{
			n.Left.Insert(k)
		}
	}
}

// Search will take in a key value
// and return true if there is a node with that value
func (n *Node) Search(k int) bool{
	if n==nil{
		return false
	}
	if k>n.Key{
		return n.Right.Search(k)
	}else if k<n.Key{
		return n.Left.Search(k)
	}
	return true
}

func main(){
	tree:=&Node{Key:100}
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(220)
	tree.Insert(210)
	tree.Insert(260)
	fmt.Println(tree)
	fmt.Println(tree.Search(260))

}