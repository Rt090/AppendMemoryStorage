package main

import (
	"bufio"
	"fmt"
	"github.com/Rt090/AppendMemoryStorage/cache"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	c := cache.NewCache()
	f, err := os.Open("./data/test1.txt")
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		l := sc.Text()
		li := strings.Split(l, ",")
		c.Insert(li[0], li[1])
	}
	c.Stats()
	out := c.Get("1")
	for _, val := range out {
		fmt.Println(val)
	}
}
