package main

import (
	"bufio"
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
	antNum := roomList.ants

	DFSSearch := AntSender(antNum, allPathsDFS)
	BFSSearch := AntSender(antNum, allPathsBFS)

	Printer := []string{}

	if len(DFSSearch) < len(BFSSearch) {
		Printer = DFSSearch
	} else {
		Printer = BFSSearch
	}

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
	firstLine := true

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		x := scanner.Text()
		fmt.Println(x)
		if firstLine {
			g.ants, _ = strconv.Atoi(x)
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
	fmt.Printf("\nPathList: %v\n", pathListStore)
	fmt.Printf("\nQUEUE: %v\n", queue)

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
	fmt.Printf("ORDER: %v", order)

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

	fmt.Printf("\n\nCONTAINER!!!!!!!!!!!!!!!!!!!!!!!!111 %v\n\n", container)
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

//Apply to ant mover function to each ant in the order list
// pathNum := 0
// allMoves := [][]string{}
// change := 1
// track := 0
// for i, num := range order {
// 	fmt.Printf("\nALL MOVES%d :  %v \n", i, allMoves)
// 	fmt.Printf("\n Track == %v || Change == %d || i == %d\n", track, change, i)

// 	if len(queue[len(queue)-change]) == track {
// 		change++
// 		fmt.Printf("\n CHANGED!!!| CHANGE = %d | i = %d | track = %d | len of PAthListSTORE = %d \n", change, i, track, len(pathListStore))
// 		track = 0
// 	} else if i%(len(pathListStore)-change+1) == 0 && i != 0 {
// 		track++
// 		fmt.Printf("\nTrack Increased!!!!!!!!!!!!! Track == %d\n", track)

// 	}

// 	if pathNum == len(pathListStore)-change+1 {
// 		fmt.Println("PATHNUM CHANGED")
// 		pathNum = 0
// 	}

// 	fmt.Printf("\nANT: %v || PathNum: %d\n", num, pathNum)
// 	allMoves = append(allMoves, AntMover(num, pathListStore[pathNum]))
// 	pathNum++
// }

// fmt.Printf("\nALL MOVES SORTED:  %v\n", allMoves)
// finalPrint := []string{}
// fmt.Printf("\nPATHLISTSTORE: %v\n", pathListStore)

// //Loop through all moves
// // add :=0
// // tracker := len(pathListStore)-1
// // var xCheck int
// for i, z := 0, 0; i < len(allMoves); i, z = i+1, z+0 {

// 	// if i >= len(queue[tracker]) * tracker+1 {
// 	// 	if add != len(queue)-1{
// 	// 		add++
// 	// 	}
// 	// 	if tracker > 0 {
// 	// 	tracker--
// 	// 	}
// 	// }
// 	// fmt.Printf("\n LEN OF PAPTHLISTSTORE: %v| Tracker = %d \n", len(queue[tracker]), tracker)
// 	fmt.Printf("\nMODULO of Len: %d \n", i%(len(pathListStore)))

// 	if i%(len(pathListStore)) == 0 && i != 0 {
// 		z++
// 	}

// 	//Loop through current element
// 	for j := 0; j < len(allMoves[i]); j++ {
// 		if z+j > len(finalPrint)-1 {
// 			finalPrint = append(finalPrint, allMoves[i][j]+" ")
// 			fmt.Printf("\nAPPENDED: %v| z= %d| j=%d| i=%d \n", allMoves[i][j], z, j, i)
// 		} else {
// 			finalPrint[z+j] += allMoves[i][j] + " "
// 			fmt.Printf("\nADDED: %v| z= %d| j=%d| i=%d\n", allMoves[i][j], z, j, i)

// 		}

// 		//Add to finalprrint at index i+j

// 	}
// 	fmt.Println(finalPrint)
// }

// fmt.Printf("\n \n \nFInal PRINT: %v", finalPrint)
// for i, t := range finalPrint {
// 	fmt.Printf("\nSTEP %d:  %v ", i, t)
// }

// fmt.Println(allMoves)
