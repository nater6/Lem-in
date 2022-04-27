package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//graph structre
// graph is adjacency list
type Graph struct {
	Rooms     []*Room
	startRoom string
	endRoom   string
	ants      int
}

// vertex represents graph vertex
type Room struct {
	Roomname string
	adjacent []string
	visited  bool
}

// add vertext
func (g *Graph) AddRoom(name string) {
	g.Rooms = append(g.Rooms, &Room{Roomname: name, adjacent: []string{}, visited: false})
}

//add edge
func (g *Graph) AddLinks(from, to string) {
	// get vertex
	fromRoom := g.getRoom(from)
	toRoom := g.getRoom(to)

	// check
	if fromRoom == nil || toRoom == nil {
		//If either of the rooms dont exsist
		err := fmt.Errorf("Room doesn't exsist (%v --- %v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromRoom.adjacent, to) || contains(toRoom.adjacent, from) {
		// if the link already exsists
		err := fmt.Errorf(" Existing Link (%v --- %v)", from, to)
		fmt.Println(err.Error())
	} else if fromRoom.Roomname == g.endRoom {
		//Checking for the endroom
		toRoom.adjacent = append(toRoom.adjacent, fromRoom.Roomname)
	} else if toRoom.Roomname == g.endRoom {
		//Checking for the endroom
		fromRoom.adjacent = append(fromRoom.adjacent, toRoom.Roomname)
	} else if toRoom.Roomname == g.startRoom {
		//Checking for the startroom
		toRoom.adjacent = append(toRoom.adjacent, fromRoom.Roomname)
	} else if fromRoom.Roomname == g.startRoom {
		//Checking for the startroom
		fromRoom.adjacent = append(fromRoom.adjacent, toRoom.Roomname)
	} else if fromRoom.Roomname != g.endRoom && toRoom.Roomname != g.endRoom {

		fromRoom.adjacent = append(fromRoom.adjacent, toRoom.Roomname)
		toRoom.adjacent = append(toRoom.adjacent, fromRoom.Roomname)

	}

}

// get vertex
func (g *Graph) getRoom(name string) *Room {
	for i, v := range g.Rooms {
		if v.Roomname == name {
			return g.Rooms[i]
		}
	}
	return nil
}

// contains
func contains(s []string, name string) bool {
	for _, v := range s {
		if name == v {
			return true
		}
	}
	return false
}

func main() {

	

	list1 := []*Room{}

	roomList := &Graph{Rooms: list1}
	

	 if err := SortFiles(roomList); err!= nil {
		fmt.Print(err) 
		return
	 }

	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan(){
		x  := scanner.Text()
		fmt.Println(x)
	}

	//Run the DFS path search
	allPathsDFS := []string{}
	allPathsBFS := []string{}
	var path string
	DFS(roomList.startRoom, roomList.endRoom, roomList, path, &allPathsDFS)

	//Run the Shortest path search
	list2 := []*Room{}
	roomList1 := &Graph{Rooms: list2}
	SortFiles(roomList1)

	BFS(roomList1.startRoom, roomList1.endRoom, roomList1, &allPathsBFS, ShortestPath)
		//Sort the path lists in order

	lenSorter(&allPathsBFS)
	lenSorter(&allPathsDFS)
	

	//Send ants using the function
	antNum := roomList.ants

	DFSSearch := AntSender(antNum, allPathsDFS)
	BFSSearch := AntSender(antNum, allPathsBFS)

	Printer := []string{}

	if len(DFSSearch) < len(BFSSearch) {
		Printer = DFSSearch
	} else {
		Printer = BFSSearch
	}
	fmt.Println()
	for _, step := range Printer {
		fmt.Println(step)
	}

}

func BFS(start, end string, g *Graph, paths *[]string, f func(graph *Graph, start string, end string, path Array) Array) {

	begin := g.getRoom(start)

	if len(begin.adjacent) == 2 {
		begin.adjacent[0], begin.adjacent[1] = begin.adjacent[1], begin.adjacent[0]
	}

	for i := 0; i < len(begin.adjacent); i++ {

		
		var shortPath Array

		//Find all possible paths with unvisited rooms
		ShortestPath(g, g.startRoom, g.endRoom, shortPath)

		// Get the value string of the shortest path
		var shortStorer string
		if len(pathArray) != 0 {
			shortStorer = pathArray[0]
		}

		for _, v := range pathArray {
			if len(v) < len(shortStorer) {
				shortStorer = v
			}
		}

		//Remove the sqr brackes form the string
		if len(pathArray) != 0 {
			shortStorer = shortStorer[1 : len(shortStorer)-1]
		}

		//Mark the rooms in the path as visited

		shortStorerSlc := strings.Split(shortStorer, " ")
		shortStorerSlc = shortStorerSlc[1:]

		//Loop through the path and mark as visited
		for z := 0; z < len(shortStorerSlc)-1; z++ {
			g.getRoom(shortStorerSlc[z]).visited = true
		}

		var pathStr string
		if len(shortStorerSlc) != 0 {
			for i := 0; i < len(shortStorerSlc); i++ {
				if i == len(shortStorerSlc)-1 {
					pathStr += shortStorerSlc[i]
				} else {
					pathStr = pathStr + shortStorerSlc[i] + "-"
				}
			}
		}

		if len(pathStr) != 0 {
			containing := false
			for _, v := range *paths {
				if v == pathStr {
					containing = true
				}
			}
			if !containing {
				*paths = append(*paths, pathStr)
			}
		}

		pathArray = []string{}
	}

}

func SortFiles(g *Graph) error {
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	start := false
	end := false
	i := 0
	firstLine := true

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		x := scanner.Text()
		if firstLine {
			g.ants, _ = strconv.Atoi(x)
			if g.ants == 0{
				return errors.New("ERROR: invalid data format")
			}
			firstLine = false
		}

		space := strings.Split(scanner.Text(), " ")

		if len(space) > 1 {
			g.AddRoom(space[0])
			i++
		}

		if start {
			g.startRoom = g.Rooms[i-1].Roomname
			start = false
		} else if end {
			g.endRoom = g.Rooms[i-1].Roomname
			end = false
		}

		hyphen := strings.Split(scanner.Text(), "-")
		if len(hyphen) > 1 {
			if hyphen[0] == hyphen[1] {
				return errors.New("ERROR: invalid data format")
				
			}
			g.AddLinks(hyphen[0], hyphen[1])

		}

		if x == "##start" {
			start = true
		}

		if x == "##end" {
			end = true
		}

	}
	return nil

}

func DFS(current, end string, g *Graph, path string, pathList *[]string) {
	//Get the pointer and all information for the current roo

	//Check if the current room is the end room
	curr := g.getRoom(current)

	if current != end {
		curr.visited = true
	}

	if curr.Roomname == g.endRoom {
		path += current
	} else if !(curr.Roomname == g.startRoom) {
		path += current + "-"
	}

	//Create bool var to to be true if the current room == end
	final := false

	if current == end {

		*pathList = append(*pathList, path)
		path = ""

		final = true

		for i := 0; i < len(g.getRoom(g.startRoom).adjacent); i++ {
			if g.getRoom(g.startRoom).adjacent[i] == g.endRoom {
				g.getRoom(g.startRoom).adjacent[i] = ""
			}
		}

	}

	if final {
		DFS(g.startRoom, end, g, path, pathList)
	}

	//Loop through adjacent rooms and see if the end room is present or if there are any unvisited rooms
	for i := 0; i < len(curr.adjacent); i++ {

		// If the end room is present in the adjacent room move it to the start of the slice
		if curr.adjacent[i] == g.endRoom {
			curr.adjacent[0], curr.adjacent[i] = curr.adjacent[i], curr.adjacent[0]
		}
	}

	// recurssively call the func to the end
	for i := 0; i < len(curr.adjacent); i++ {
		if curr.adjacent[i] == "" {
			continue
		}
		//Get information for the current room
		x := g.getRoom(curr.adjacent[i])

		if x.visited {
			continue
		} else {
			DFS(x.Roomname, end, g, path, pathList)
		}
	}
}

type Array []string

var pathArray Array

func (arr Array) hasPropertyOf(str string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}

func ShortestPath(graph *Graph, start string, end string, path Array) Array {
	path = append(path, start)
	if start == end {
		return path
	}
	shortest := make([]string, 0)
	for _, node := range graph.getRoom(start).adjacent {
		if !path.hasPropertyOf(node) && !graph.isVisited(node) {
			newPath := ShortestPath(graph, node, end, path)
			if len(newPath) > 0 {
				if newPath.hasPropertyOf(graph.startRoom) && newPath.hasPropertyOf(end) {
					pathArray = append(pathArray, fmt.Sprint(newPath))
					if len(shortest) == 0 || (len(newPath) < len(shortest)) {

						shortest = newPath
					}
				}
			}
		}
	}
	return shortest
}

func (graph *Graph) isVisited(str string) bool {
	return graph.getRoom(str).visited
}

func lenSorter(paths *[]string) {

	x := *paths
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x); j++ {
			if len(x[i]) < len(x[j]) {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
	*paths = x
}

func AntSender(n int, pathList []string) []string {
	pathListStore := [][]string{}

	for _, v := range pathList {
		s := strings.Split(v, "-")
		pathListStore = append(pathListStore, s)
	}

	lenP := len(pathList)

	queue := make([][]string, lenP)

	x := 0

	for i := 1; i <= n; i++ {
		ant := strconv.Itoa(i)

		if x == lenP-1 {
			if len(pathListStore[x])+len(queue[x]) <= len(pathListStore[0])+len(queue[0]) {
				queue[x] = append(queue[x], ant)
			} else {
				x = 0
				queue[x] = append(queue[x], ant)
			}

		} else {
			if len(pathListStore[x])+len(queue[x]) <= len(pathListStore[x+1])+len(queue[x+1]) {
				queue[x] = append(queue[x], ant)
			} else {
				x++
				queue[x] = append(queue[x], ant)

			}
		}
	}

	longest := len(queue[0])

	for i := 0; i < len(queue); i++ {
		if len(queue[i]) > longest {
			longest = len(queue[i])
		}
	}

	order := []int{}

	for j := 0; j < longest; j++ {
		for i := 0; i < len(queue); i++ {
			if j < len(queue[i]) {
				x, _ = strconv.Atoi(queue[i][j])
				order = append(order, x)
			}
		}
	}

	container := make([][][]string, len(queue))

	for i := 0; i < len(queue); i++ {

		for _, a := range queue[i] {
			adder := []string{}
			for _, room := range pathListStore[i] {
				str := "L" + a + "-" + room
				adder = append(adder, str)
			}
			container[i] = append(container[i], adder)

		}
	}
	finalMoves := []string{}

	for _, paths := range container {
		for j, moves := range paths {
			for k, room := range moves {
				if j+k > len(finalMoves)-1 {
					finalMoves = append(finalMoves, room+" ")
				} else {
					finalMoves[j+k] += room + " "
				}
			}

		}

	}

	return finalMoves

}

func AntMover(n int, path []string) []string {
	antRooms := []string{}
	x := strconv.Itoa(n)
	for i := 0; i < len(path); i++ {
		str := "L" + x + "-" + path[i]
		antRooms = append(antRooms, str)
	}
	return antRooms
}
