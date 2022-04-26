package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//graph structre
// graph is adjacency list
type Graph struct {
	Rooms     []*Room
	startRoom string
	endRoom   string
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

	SortFiles(roomList)

	for _, r := range roomList.Rooms {
		fmt.Println(r.Roomname)
		fmt.Println(r.adjacent)

	}

	for _, v := range roomList.Rooms {
		x := *v
		fmt.Print(x.Roomname + "|")

	}
	fmt.Println()
	fmt.Println("startroom: " + roomList.startRoom)
	fmt.Println("endroom: " + roomList.endRoom)

	//Run the DFS path search
	allPathsDFS := []string{}
	allPathsBFS := []string{}
	var path string
	DFS(roomList.startRoom, roomList.endRoom, roomList, path, &allPathsDFS)
	fmt.Println(allPathsDFS)

	//Run the Shortest path search
	list2 := []*Room{}
	roomList1 := &Graph{Rooms: list2}
	SortFiles(roomList1)

	BFS(roomList1.startRoom, roomList1.endRoom, roomList1, &allPathsBFS, ShortestPath)
	fmt.Println(allPathsBFS)

	//Sort the path lists in order

	lenSorter(&allPathsBFS)
	lenSorter(&allPathsDFS)
	fmt.Printf("\nPRINTING DFS PATHS: %v\n", allPathsDFS)
	fmt.Printf("\nPRINTING BFS PATHS: %v\n", allPathsBFS)

	//Send ants using the function

}

func BFS(start, end string, g *Graph, paths *[]string, f func(graph *Graph, start string, end string, path Array) Array) {

	begin := g.getRoom(start)

	if len(begin.adjacent) == 2 {
		begin.adjacent[0], begin.adjacent[1] = begin.adjacent[1], begin.adjacent[0]
	}

	for i := 0; i < len(begin.adjacent); i++ {

		fmt.Printf("ROUND NUMBER: %d", i)
		var shortPath Array

		//Find all possible paths with unvisited rooms
		ShortestPath(g, g.startRoom, g.endRoom, shortPath)

		// Get the value string of the shortest path
		fmt.Printf("PATH ARRAY!!!!: %v \n", pathArray)
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

		fmt.Printf("\n                   ShortStORERSSS: %v \n", shortStorerSlc)
		//Loop through the path and mark as visited
		for z := 0; z < len(shortStorerSlc)-1; z++ {
			g.getRoom(shortStorerSlc[z]).visited = true
		}
		fmt.Printf("PATH TO BE APPENDED: %v \n", shortStorerSlc)

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

func SortFiles(g *Graph) {
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	start := false
	end := false
	i := 0

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		x := scanner.Text()
		fmt.Println(x)

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
			g.AddLinks(hyphen[0], hyphen[1])

		}

		if x == "##start" {
			start = true
			fmt.Println("Start")

		}

		if x == "##end" {
			end = true
		}

	}

}

func DFS(current, end string, g *Graph, path string, pathList *[]string) {
	//Get the pointer and all information for the current room
	fmt.Println("Current Room: " + current)

	//Check if the current room is the end room
	curr := g.getRoom(current)
	fmt.Printf("current adj list: %v", curr.adjacent)

	if current != end {
		curr.visited = true
	}

	if curr.Roomname == g.endRoom {
		path += current
	} else if !(curr.Roomname == g.startRoom) {
		path += current + "-"
	}

	fmt.Println("Path: " + path)

	//Create bool var to to be true if the current room == end
	final := false

	if current == end {
		fmt.Printf("Path 123: %v \n", path)

		*pathList = append(*pathList, path)
		fmt.Printf("appended PathList: %v \n", pathList)
		path = ""

		final = true

		for i := 0; i < len(g.getRoom(g.startRoom).adjacent); i++ {
			fmt.Println("Current value: " + g.getRoom(g.startRoom).adjacent[i])
			if g.getRoom(g.startRoom).adjacent[i] == g.endRoom {
				g.getRoom(g.startRoom).adjacent[i] = ""
				fmt.Println("End Removed")
			}
		}
		fmt.Println(g.getRoom(g.startRoom).adjacent)

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
		fmt.Println("Current adjacent Room: " + x.Roomname)

		if x.visited {
			fmt.Println("Previously visited: " + x.Roomname)
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
					fmt.Printf("\n New Path: %v \n", newPath)
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