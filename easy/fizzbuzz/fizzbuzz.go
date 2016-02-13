package main

import (
    "fmt"
    "os"
    "bufio"
	"strconv"
    "strings"
)

func parseInts(s []string) []int64 {
	var ints []int64
	for _,s := range s {
		x, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		} else {
			ints = append(ints,x)
		}
	}
	return ints
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	var x,y,n int64
	var ai []int64
    io := bufio.NewReader(f)//(os.Stdin)
	for {
		s, err := io.ReadString('\n') // split on new line
		if err != nil {
			break
		}
		ai = parseInts(strings.Split(strings.Trim(s,"\n")," ")) // parse the ints
		x = ai[0]
		y = ai[1]
		n = ai[2]
		var i int64
		for i = 1; i<=n; i++ {
			if i%x==0 && i%y==0 {
				fmt.Printf("FB ")
			} else if i%x==0 {
				fmt.Printf("F ")
			} else if i%y==0 {
				fmt.Printf("B ")
			} else {
				fmt.Printf("%d ",i)
			}
		}
		fmt.Printf("\n")
	}
}
