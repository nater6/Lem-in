package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// The Graph structure keeps track of all rooms the ant can take, the start and end rooms of the path and the number of ants
type Graph struct {
	Rooms     []*Room
	startRoom string
	endRoom   string
	ants      int
}

// The Room structure keeps track of the roomname, The rooms that the the current room is connected to and if the room has been visited before
type Room struct {
	Roomname string
	adjacent []string
	visited  bool
}

// AddRoom is a method that adds a new room, name, to a graph
func (g *Graph) AddRoom(name string) {
	g.Rooms = append(g.Rooms, &Room{Roomname: name, adjacent: []string{}, visited: false})
}

//AddLinks is a method that adds a link from one room to another
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
		//Checking for the endroom, if the endroom is present only add the link towards the endroom and not the otherway
		toRoom.adjacent = append(toRoom.adjacent, fromRoom.Roomname)
	} else if toRoom.Roomname == g.endRoom {
		//Checking for the endroom, if the endroom is present only add the link towards the endroom and not the otherway
		fromRoom.adjacent = append(fromRoom.adjacent, toRoom.Roomname)
	} else if toRoom.Roomname == g.startRoom {
		//Checking for the startroom, if the startroom is present only add the link towards the startroom and not the otherway
		toRoom.adjacent = append(toRoom.adjacent, fromRoom.Roomname)
	} else if fromRoom.Roomname == g.startRoom {
		//Checking for the startroom, if the startroom is present only add the link towards the startroom and not the otherway
		fromRoom.adjacent = append(fromRoom.adjacent, toRoom.Roomname)
	} else if fromRoom.Roomname != g.endRoom && toRoom.Roomname != g.endRoom {
		//If both rooms are not endrooms then add the link in both directions
		fromRoom.adjacent = append(fromRoom.adjacent, toRoom.Roomname)
		toRoom.adjacent = append(toRoom.adjacent, fromRoom.Roomname)

	}

}

// getRoom is a method that returns a pointer to the room 'name'
func (g *Graph) getRoom(name string) *Room {
	for i, v := range g.Rooms {
		if v.Roomname == name {
			return g.Rooms[i]
		}
	}
	return nil
}

// Contains check whether the string name is present in the slice of strings s
func contains(s []string, name string) bool {
	for _, v := range s {
		if name == v {
			return true
		}
	}
	return false
}

func main() {
	//create a graph
	list1 := []*Room{}
	roomList := &Graph{Rooms: list1}

	//sort thorugh the .txt file
	if err := SortFiles(roomList); err != nil {
		fmt.Print(err)
		return
	}

	//Print the contents of the file in the terminal
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		x := scanner.Text()
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

	// Preform the Searches using BFS and DFS 
	DFSSearch := AntSender(antNum, allPathsDFS)
	BFSSearch := AntSender(antNum, allPathsBFS)

	Printer := []string{}

	//compare the DFS and BFS search and print the shorter one
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

//BFS preforms a Breadth First Search of a graph from rooms start to end and puts all paths found in the []string paths
func BFS(start, end string, g *Graph, paths *[]string, f func(graph *Graph, start string, end string, path Array) Array) {
	//begin is the pointer to the start room of the search
	begin := g.getRoom(start)

	//if (len==2) the second adjacent room is the end room because of the way the rooms are added so we switch them to make the endroom the first adjacent room checked
	if len(begin.adjacent) == 2 {
		begin.adjacent[0], begin.adjacent[1] = begin.adjacent[1], begin.adjacent[0]
	}

	//loop through the adjacent rooms
	for i := 0; i < len(begin.adjacent); i++ {

		//shortpath is a variable we will use the path we are checking (Array is a structure of type []string made earlier)
		var shortPath Array

		//Shortest path is a function used to find the possible paths from start to end, and sort them by length in ascending order, using unvisited room
		ShortestPath(g, g.startRoom, g.endRoom, shortPath)

		// shortstorer is a variable that stores the shortest path from the start to the end
		var shortStorer string
		if len(pathArray) != 0 {
			shortStorer = pathArray[0]
		}

		for _, v := range pathArray {
			if len(v) < len(shortStorer) {
				shortStorer = v
			}
		}

		//Remove the square brackes form the string
		if len(pathArray) != 0 {
			shortStorer = shortStorer[1 : len(shortStorer)-1]
		}

		//Turn the path string into a slice to mark in as visited

		shortStorerSlc := strings.Split(shortStorer, " ")
		shortStorerSlc = shortStorerSlc[1:]

		//Loop through the path and mark as visited
		for z := 0; z < len(shortStorerSlc)-1; z++ {
			g.getRoom(shortStorerSlc[z]).visited = true
		}

		//loop throigh the slice and turn it back into a string
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

		//Check if the path found is already in the the overall path []string, if its not then add it
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
		//Reset pathArray to check next adjacent room if no path found
		pathArray = []string{}
	}

}

//SortFiles sorts goes through the txt file with the ants and rooms and adds the rooms and links to the graph
func SortFiles(g *Graph) error {
	//Create a scanner to go through each line of the txt file
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	start := false
	end := false
	i := 0
	firstLine := true

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		x := scanner.Text()
		//Check if there is a valid number of ants on the file
		if firstLine {
			g.ants, _ = strconv.Atoi(x)
			if g.ants == 0 {
				return errors.New("ERROR: invalid data format")
			}
			firstLine = false
		}

		//create a []string that holds all information in the current line of the text file
		space := strings.Split(scanner.Text(), " ")

		//If len > 1 the line is specifiying a room with the first element of the slice being the room, use Addroom to add the room the the graph
		if len(space) > 1 {
			g.AddRoom(space[0])
			i++
		}

		//Check if the room is the start or end room if so assign the start/endroom element of the graph to the room
		if start {
			g.startRoom = g.Rooms[i-1].Roomname
			start = false
		} else if end {
			g.endRoom = g.Rooms[i-1].Roomname
			end = false
		}

		//Now check if the line is specifying links between rooms
		hyphen := strings.Split(scanner.Text(), "-")
		//If the length > 1 then the line of the txt file is linking two rooms together
		if len(hyphen) > 1 {
			//Check if both rooms on the line are the same (If the room links to itself)
			if hyphen[0] == hyphen[1] {
				return errors.New("ERROR: invalid data format")

			}
			//If they dont link to them selves use the AddLinks method to add the link to both room
			g.AddLinks(hyphen[0], hyphen[1])

		}

		//Check if the line after will contain the start/endroom
		if x == "##start" {
			start = true
		}

		if x == "##end" {
			end = true
		}

	}
	//If the txt file is a valid format return no error
	return nil

}

//DFS preforms a depth first search of a graph and returns the possible paths
func DFS(current, end string, g *Graph, path string, pathList *[]string) {

	//Check if the current room is the end room
	curr := g.getRoom(current)

	//Mark the room as visited (Unless it is the end room, so the endroom can be accessed in another path)
	if current != end {
		curr.visited = true
	}

	//Add the current room to the path string
	if curr.Roomname == g.endRoom {
		path += current
	} else if !(curr.Roomname == g.startRoom) {
		path += current + "-"
	}

	//Create bool var to to be true if the current room == end
	final := false
	//Check if the current room is the end room
	if current == end {
		//Add the path found to the pathroom
		*pathList = append(*pathList, path)
		//Reset the path to find another room
		path = ""

		final = true

		//If the start room is adjacent to the endroom remove the link as it would be found first
		for i := 0; i < len(g.getRoom(g.startRoom).adjacent); i++ {
			if g.getRoom(g.startRoom).adjacent[i] == g.endRoom {
				g.getRoom(g.startRoom).adjacent[i] = ""
			}
		}

	}

	//If the end room has been found call the function again, from the the start room to check for another path
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
		//if the element in the adj room is "" then its empty so skip
		if curr.adjacent[i] == "" {
			continue
		}
		//Get information for the current room
		x := g.getRoom(curr.adjacent[i])

		//If the room has been visited before skip, if not check the next recursivly call the DFS func from the next room
		if x.visited {
			continue
		} else {
			DFS(x.Roomname, end, g, path, pathList)
		}
	}
}

type Array []string

var pathArray Array

//hasPropertyOf is a method that checks if an array/[]string contains a string
func (arr Array) hasPropertyOf(str string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}

//ShortestPath finds all the possible paths from start to end room using BFS and sorts them in ascending order
func ShortestPath(graph *Graph, start string, end string, path Array) Array {
	//Add the room to the [ath]
	path = append(path, start)
	//If the currentroom is the endroom retrun the path
	if start == end {
		return path
	}

	shortest := make([]string, 0)
	//go through the adjacent rooms to the start room
	for _, node := range graph.getRoom(start).adjacent {
		//if the current path doesn't contain the adj room and the room is not visited recursively call the function dhortest path
		if !path.hasPropertyOf(node) && !graph.isVisited(node) {
			newPath := ShortestPath(graph, node, end, path)
			if len(newPath) > 0 {
				if newPath.hasPropertyOf(graph.startRoom) && newPath.hasPropertyOf(end) {
					pathArray = append(pathArray, fmt.Sprint(newPath))
					// if len(shortest) == 0 || (len(newPath) < len(shortest)) {

					// 	shortest = newPath
					// }
				}
			}
		}
	}
	return shortest
}

func (graph *Graph) isVisited(str string) bool {
	return graph.getRoom(str).visited
}


//lenSorter sorts all paths found in ascending order
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


//AntSender puts the 
func AntSender(n int, pathList []string) []string {
	pathListStore := [][]string{}
	

	//go through each path and turn the string into []string where each element is a room on the path
	for _, v := range pathList {
		s := strings.Split(v, "-")
		pathListStore = append(pathListStore, s)
	}


	lenP := len(pathList)

	//queue is a [][]strings with the ants going down each path in each slice [[1234][567][89]]
	queue := make([][]string, lenP)

	x := 0

	for i := 1; i <= n; i++ {
		ant := strconv.Itoa(i)
		//If the number of steps in the curr room + num of ants is less than that of the next add to the current room else add the the next (if x is at the last room lenP-1 set x back to the first room)
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

	//loop through all the queue and determine the order 

	for j := 0; j < longest; j++ {
		for i := 0; i < len(queue); i++ {
			if j < len(queue[i]) {
				//Adds the ants from each path that can run at the same time
				x, _ = strconv.Atoi(queue[i][j])
				order = append(order, x)
			}
		}
	}

	//container is a [][][]string that holds all the movements
	container := make([][][]string, len(queue))

	//Loop thorugh the queue and send the ants
	for i := 0; i < len(queue); i++ {
		//loop through each room of each path and add each ants movements
		for _, a := range queue[i] {
			adder := []string{}
			for _, room := range pathListStore[i] {
				str := "L" + a + "-" + room
				adder = append(adder, str)
			}
			//add all of the rooms of ant a to the container
			container[i] = append(container[i], adder)

		}
	}
	//Final moves is a []string where each element is one step of the result
	finalMoves := []string{}

	//loop thorugh container and add all the moves to one slice of strings
	for _, paths := range container {
		for j, moves := range paths {
			for k, room := range moves {
				//if finalmoves doesnt have aan moving at that index yet add the index
				if j+k > len(finalMoves)-1 {
					finalMoves = append(finalMoves, room+" ")
				} else {
					//if an ant is moving at a step that has already been added add it to the step
					finalMoves[j+k] += room + " "
				}
			}
		}
	}
	return finalMoves
}

// func AntMover(n int, path []string) []string {
// 	antRooms := []string{}
// 	x := strconv.Itoa(n)
// 	for i := 0; i < len(path); i++ {
// 		str := "L" + x + "-" + path[i]
// 		antRooms = append(antRooms, str)
// 	}
// 	return antRooms
// }


