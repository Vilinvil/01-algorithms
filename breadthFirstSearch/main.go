package main

import (
	"container/list"
	"fmt"
	"strings"
)

// Linear search element in slice.
func isElemInSlice(elem string, sl []string) bool {
	for _, v := range sl {
		if v == elem {
			return true
		}
	}
	return false
}

// Goal this program is find killer among friends
func main() {
	// Create queue based on doubly linked list. (FIFO)
	queue := list.New()

	// Init graph friends of man.
	graph := make(map[string][]string)
	graph["I"] = []string{"Ivan manager", "Danil engineer", "Alex itGay"}
	graph["Ivan manager"] = []string{"Nik student", "Ken mechanic"}
	graph["Danil engineer"] = []string{"Dina office worker ", "Lena assistant manager"}
	graph["Alex itGay"] = []string{"Natalia marketer", "Bob killer"}

	trustPeople := make([]string, 0) // Slice for people who can be trusted

	queue.PushBack("I")
	for queue.Len() > 0 {
		el := queue.Front()
		//str, err := el.Value.(string)
		str := fmt.Sprintf("%v", el.Value)
		//if err == true {
		//	log.Fatalf("Element %v don`t convert to string ", el)
		//}
		//fmt.Println(str)
		if isElemInSlice(str, trustPeople) || !strings.Contains(str, "killer") { //killer check
			//fmt.Println("inside")
			for _, val := range graph[str] {
				queue.PushBack(val)
				//fmt.Println("inside 2")
			}
			queue.Remove(el)
			continue
		}
		fmt.Printf("%v is killer!!!\n", str)
		break
	}
}
