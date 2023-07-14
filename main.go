package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GetRandom(num int) int {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	randomInt := r.Intn(num)
	// fmt.Println(randomInt)
	return randomInt
}

func jj() {
	a := GetRandom(1000)
	b := GetRandom(1000)
	c := GetRandom(2)
	y := "+"
	if c == 0 {
		y = "-"
	}
	// fmt.Println(c)
	if a > b {
		fmt.Println(strconv.Itoa(a) + y + strconv.Itoa(b) + " = ")
	} else if a < b {
		fmt.Println(strconv.Itoa(b) + y + strconv.Itoa(a) + " = ")
	} else {
		jj()
	}
}

func main() {
	fmt.Println("==加减口算==")
	for i := 0; i < 10; i++ {
		jj()
	}
	fmt.Println("==口算==")
}
