package main

import (
    "github.com/RyanCarrier/dijkstra"
    "fmt"
)

func main() {
	graph := dijkstra.NewGraph()

	graph.AddVertex(0)
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddVertex(5)
	graph.AddVertex(6)
	graph.AddVertex(7)
	graph.AddVertex(8)
	graph.AddVertex(9)

	graph.AddArc(0, 3, 12)
	graph.AddArc(0, 5, 4)

    graph.AddArc(1, 2, 13)
    graph.AddArc(1, 4, 1)
    graph.AddArc(1, 3, 21)

	graph.AddArc(2, 0, 2)

    graph.AddArc(3, 5, 1)

    graph.AddArc(4, 0, 7)

    graph.AddArc(5, 8, 8)
    graph.AddArc(5, 9, 3)
    
    graph.AddArc(6, 0, 1)
    graph.AddArc(6, 4, 2)
    graph.AddArc(6, 7, 12)
    graph.AddArc(6, 8, 3)
    
    graph.AddArc(7, 0, 3)
    graph.AddArc(7, 9, 2)
    graph.AddArc(7, 8, 21)
    
    graph.AddArc(9, 5, 4)
    graph.AddArc(9, 8, 4)
    graph.AddArc(9, 6, 1)

    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    fmt.Println("Executing...")
    for _, num := range nums {
        best, err := graph.Shortest(0,num)
        if err!=nil{
            fmt.Println(err, " from 0 to ", num)
        }else{
            fmt.Println("Shortest distance from 0 to ", num, ". Total distance: ", best.Distance, " following path ", best.Path)
        }
    }

}
