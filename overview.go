// Hash Map - Hash Set
// LinkedList
// Tree/Graph/DP
// Array/String/Sequence
// Stream

// Search(filter),Sort,Map/Reduce

package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go fmt.Println("Hello ", i)
	}
	time.Sleep(1 * time.Second)
}
