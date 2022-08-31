package main

import (
	"container/list"
	"fmt"
	"log"
	"strings"
)

// IsElemInSlice is func linear search element in slice.
func IsElemInSlice(elem string, sl []string) bool {
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

	// Init graph friends of man. This graph is singly connected(односвязный)
	graph := make(map[string][]string)
	graph["I"] = []string{"Ivan manager", "Danil engineer", "Alex itGay"}
	graph["Ivan manager"] = []string{"Nik student", "Ken mechanic"}
	graph["Danil engineer"] = []string{"Dina office worker ", "Lena assistant manager"}
	graph["Alex itGay"] = []string{"Natalia marketer", "Bob killer"}

	trustPeople := make([]string, 0) // Slice for people who can be trusted

	queue.PushBack("I")
	for queue.Len() > 0 {
		el := queue.Front()
		str, ok := el.Value.(string)
		if !ok {
			log.Fatalf("Element %v don`t convert to string ", el)
		}

		if !IsElemInSlice(str, trustPeople) {
			if strings.Contains(str, "killer") { //killer check
				fmt.Printf("%v is killer!!!\n", str)
				break
			}

			for _, val := range graph[str] {
				queue.PushBack(val)
			}

			trustPeople = append(trustPeople, str)
			queue.Remove(el)
			continue
		}
	}
}
