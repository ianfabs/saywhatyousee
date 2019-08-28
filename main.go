package main

import "fmt"
import "strings"
import "strconv"

func main() {
	fmt.Println(swys("1", 8))
	// fmt.Println(strings.Fields("1 11"))
}

func swys(n string, max int) string {
	fields := strings.Fields(n)
	if len(fields) == max {
		return n
	}
	last := fields[len(fields)-1]
	counts := make(map[string]int)
	var ch string
	var i int
	var size int
	for size = len(last); i < size; i++ {
		ch = string(last[i])
		if counts[ch] != 0 {
			counts[ch] = counts[ch] + 1
		} else {
			counts[ch] = 1
		}
	}
	var str string
	for key, val := range counts {
		str += strconv.Itoa(val) + key
	}
	return swys(n+" "+str, max)
}
