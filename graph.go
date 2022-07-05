package main

import "fmt"

// Graph is a abstract data type and can be represented in two ways
// 1- Adjacency list (less space complexity)
// 2- Adjacency Matrix (less time complexity, need more space)

// Here we are implementing graph in Adjacency list

// Graph represents an adjacency list graph
type Graph struct{
	vertices []*Vertex
}

// Vertex represens a graph vertex
type Vertex struct{
	key int
	adjacent []*Vertex
}

// AddVertex adds a Vertex to the Graph
func (g *Graph) AddVertex(k int){
	if contains(g.vertices,k){
		err:=fmt.Errorf("Vertex %v alredy exist",k)
		fmt.Println(err.Error())
	}else{
		g.vertices=append(g.vertices,&Vertex{key:k})
	}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from int,to int){
	// get vertex
	fromVertex:=g.getVertex(from)
	toVertex:=g.getVertex(to)
	// check error
	if fromVertex==nil || toVertex==nil{
		err:=fmt.Errorf("Invalid edge (%v-->%v)",from,to)
		fmt.Println(err.Error())
	}else if contains(fromVertex.adjacent,to){
		err:=fmt.Errorf("Existing edge (%v-->%v)",from,to)
		fmt.Println(err.Error())
	}else{
	// add edge
	fromVertex.adjacent=append(fromVertex.adjacent,toVertex)
	}
}

// getVertex returns a pointer to the vertex with a key integer
func (g *Graph) getVertex(k int) *Vertex{
	for i,v:=range g.vertices{
		if k==v.key{
			return g.vertices[i]
		}
	}
	return nil
}

// Contains
func contains(s []*Vertex,k int) bool{
	for _,v:=range s{
		if k==v.key{
			return true
		}
	}
	return false
}

// Print will print the adjascent lit or eac vertex of the graph
func (g *Graph) Print(){
	for _,v:=range g.vertices{
		fmt.Printf("\nVertex %v : ",v.key)
		for _,v:=range v.adjacent{
			fmt.Printf(" %v ",v.key)
		}		
		}
		fmt.Println()
}	
func main(){
	test:=&Graph{}
	
	for i:=0;i<5;i++{
		test.AddVertex(i)
	}
	test.AddEdge(1,2)
	test.AddEdge(6,2)
	test.AddEdge(3,2)
	test.AddEdge(1,2)

	test.Print()
}