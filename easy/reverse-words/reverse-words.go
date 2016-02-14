package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func main() {
	var as []string
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	io := bufio.NewReader(f)
	for {
		s, err := io.ReadString('\n')
		if err != nil {break}
		as = strings.Split(strings.Trim(s,"\n")," ")
		for i := (len(as)-1); i>=0; i-- {
			fmt.Printf("%s ", as[i])
		}
		fmt.Println()
	}
}
