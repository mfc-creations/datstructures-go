package main

import "fmt"

// AlphabetSize is number of possible characters in the trie
const AlphabetSize=26

// Node represents each node in the trie
type Node struct{
	children [AlphabetSize]*Node
	isEnd bool
}

// Trie represent a trie and has a pointer to the root node
type Trie struct{
	root *Node
}

// Insert
func(t *Trie) Insert(w string){
	wordLength:=len(w)
	currentNode:=t.root
	for i:=0;i<wordLength;i++{
		charIndex:=w[i]-'a'  // 'a'=97
		if currentNode.children[charIndex]==nil{
			currentNode.children[charIndex]=&Node{}
		}
		currentNode=currentNode.children[charIndex]
	}
	currentNode.isEnd=true
}

// Sarch will take in a word and return true is that word is included in the tri
func (t *Trie) Search(w string) bool{
	wordLength:=len(w)
	currentNode:=t.root
	for i:=0;i<wordLength;i++{
		charIndex:=w[i]-'a'
		if currentNode.children[charIndex]==nil{
			return false
		}
		currentNode=currentNode.children[charIndex]
	}
	if currentNode.isEnd==true{
		return true
	}
	return false
}

// InitTrie will crete new trie
func InitTrie() *Trie{
	result:=&Trie{root:&Node{}}
	return result
}



// func main(){
// 	myTrie:=InitTrie()
	
// 	words:=[]string{
// 		"oreo",
// 		"mfc",
// 		"company",
// 		"golang",
// 		"dsa",
// 	}

// 	for _,v :=range words{
// 		myTrie.Insert(v)
// 	}

// 	fmt.Println(myTrie.Search("mfc"))
// 	fmt.Println(myTrie.Search("code"))
// 	fmt.Println(myTrie.Search("dsa"))
// }