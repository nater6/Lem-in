package main

import (
	"bufio"
	"fmt"
	"lemin/lemin"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Roomname string
	x        float64
	y        float64
}

type Link struct {
	startroom string
	endroom   string
	distance  float64
}

type Path struct {
	start bool
	final bool
}
//graph structre 
// graph is adjacency list
type Graph struct {
	vertices []*Vertex

}

// vertex represents graph vertex
type Vertex struct{
	key int
	adjacent []*Vertex
}
// add vertext
func (g *Graph) AddVertex(k int) {
	g.vertices = append(g.vertices, &Vertex{key: k})
}
//add edge

func (g *Graph) AddEdge(from,to int) {
	// get vertex 
	fromVertex := g.getVertex(from)
	toVertex:= g.getVertex(to)

// check error
if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v ---> %v)", from, to)
		fmt.Println(err.Error())
} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf(" Exsisting edge (%v ---> %v)", from, to)
		fmt.Println(err.Error())
	} else {
	//add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
}
	


}
// get vertex
func (g *Graph) getVertex(k int) *Vertex {
	for i,v:= range g.vertices {
		if v.key == k{
			return g.vertices[i]
		}
	}
	return nil
}
// contains
func contains(s []*Vertex, k int) bool {
			for _,v := range s {
				if k== v.key {
					return true 
				}
			}
			return false 
}
// print will print the adjacent list for each vertex of the graph 
func (g *Graph) Print() {
	for _, v := range g.vertices {
			fmt.Printf("\nVertex %v: ", v.key)
			for _, v := range v.adjacent {
				fmt.Printf("%v", v.key)
			}
	}
}

func main() {
	Rooms, Links := SortFiles()
	Links = FindDistances(Rooms, Links)
	
	test := &Graph{}

	for i:= 0; i<5; i++ {
		test.AddVertex(i)
	}
	test.Print()
	test.AddEdge(1,2)
}

func SortFiles() ([]Room, []Link){
	file, _:= os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	
	scanner.Split(bufio.ScanLines)
	
	var Rooms []Room
 var Links []Link
	for scanner.Scan(){
		space := strings.Split(scanner.Text(), " ")
		if len(space) > 1 {
			//Roomname
			newx, _:= strconv.ParseFloat(space[1], 64)
			newy, _:= strconv.ParseFloat(space[2], 64)
			NewRoom := Room {Roomname: space[0], x: newx, y: newy}
			Rooms = append(Rooms, NewRoom)
		} 
		hyphen := strings.Split(scanner.Text(), "-")

		if len(hyphen) > 1 {
			//Link
			NewLink := Link{startroom:hyphen[0], endroom:hyphen[1]}
			Links = append(Links, NewLink)
		} 
	}

	return Rooms, Links
}

func FindDistances(Rooms []Room, Links []Link) []Link{
	var x1 float64
	var y1 float64
	var x2 float64
	var y2 float64

for i,l := range Links{
		for _,r := range Rooms {
			if r.Roomname == l.startroom {
				//x1, x2
				x1 = r.x
				y1 = r.y
			}
			if r.Roomname == l.endroom {
				//x2, y2
				x2 = r.x
				y2 = r.y
			}
		}
		distance:= lemin.DistanceCalc(x1,y1,x2,y2)
		p:= &Links[i]
		p.distance = distance
	}
		
	return Links
}
