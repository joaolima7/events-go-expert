package main

import "fmt"

func main() {
	event := []string{"event1", "event2", "event3"}
	event = append(event[:0], event[1:]...)
	fmt.Println(event)
}
