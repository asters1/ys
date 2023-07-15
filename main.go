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

func jian() {
	a := GetRandom(1000)
	b := GetRandom(1000)
	// fmt.Println(c)
	if a > b {
		fmt.Println(strconv.Itoa(a) + "-" + strconv.Itoa(b) + " = ")
	} else if a < b {
		fmt.Println(strconv.Itoa(b) + "-" + strconv.Itoa(a) + " = ")
	} else {
		jian()
	}
}

func jia() {
	a := GetRandom(500)
	b := GetRandom(500)
	// fmt.Println(c)
	fmt.Println(strconv.Itoa(a) + "+" + strconv.Itoa(b) + " = ")
}

func chen() {
	a := 0
	b := 0
	for a < 10 {
		a = GetRandom(20)
		// fmt.Println(a)
	}
	for b < 2 {
		b = GetRandom(10)
	}
	fmt.Println(strconv.Itoa(a) + "X" + strconv.Itoa(b) + " = ")
}

func chu() {
	a := 0
	b := 0
	for a < 10 {
		a = GetRandom(100)
		// fmt.Println(a)
	}
	for b < 2 {
		b = GetRandom(10)
	}
	fmt.Println(strconv.Itoa(a) + "÷" + strconv.Itoa(b) + " = ")
}

// ==元角分===
// 分到元
func fentoyuan() {
	a := 0
	for a < 100 {
		a = GetRandom(1000)
		// fmt.Println(a)
	}
	fmt.Println(strconv.Itoa(a) + "分=()角()分=()元()角()分")
}

// 角到元
func jiaotoyuan() {
	a := 0
	for a < 10 {
		a = GetRandom(100)
		// fmt.Println(a)
	}
	fmt.Println(strconv.Itoa(a) + "角=()元()角")
}

func yuanjiaofen() {
	fentoyuan()
	fentoyuan()
	jiaotoyuan()
	jiaotoyuan()
}

// 时分秒
func fentoshi() {
	a := 0
	for a < 60 {
		a = GetRandom(360)
		// fmt.Println(a)
	}

	fmt.Println(strconv.Itoa(a) + "分=()时()分")
}

func miaotofen() {
	a := 0
	for a < 60 {
		a = GetRandom(360)
		// fmt.Println(a)
	}

	fmt.Println(strconv.Itoa(a) + "秒=()分()秒")
}

func shifenmiaoyuansuan() {
	a := 0
	for a < 1 {
		a = GetRandom(12)
		// fmt.Println(a)
	}
	b := 0
	for b < 1 {
		b = GetRandom(60)
		// fmt.Println(a)
	}
	c := 0
	for c < 1 {
		c = GetRandom(12)
		// fmt.Println(a)
	}
	d := 0
	for d < 1 {
		d = GetRandom(60)
		// fmt.Println(a)
	}

	aa := strconv.Itoa(a) + "时" + strconv.Itoa(b) + "分"
	bb := strconv.Itoa(c) + "时" + strconv.Itoa(d) + "分"

	fmt.Println(aa + "+" + bb + "=")
}

func shifenmiao() {
	fentoshi()
	miaotofen()
	shifenmiaoyuansuan()
}

// 长度单位计算
func mizhuanhua() {
	a := 0
	for a < 1000 {
		a = GetRandom(10000)
		// fmt.Println(a)
	}
	fmt.Println(strconv.Itoa(a) + "毫米=()米()分米()厘米()毫米")
}

func changduyunsuan() {
	a := 0
	for a < 1 {
		a = GetRandom(10)
		// fmt.Println(a)
	}
	b := 0
	for b < 10 {
		b = GetRandom(100)
		// fmt.Println(a)
	}
	c := 0
	for c < 1 {
		c = GetRandom(10)
		// fmt.Println(a)
	}
	d := 0
	for d < 10 {
		d = GetRandom(100)
		// fmt.Println(a)
	}
	aa := strconv.Itoa(a) + "米" + strconv.Itoa(b) + "厘米"
	bb := strconv.Itoa(c) + "米" + strconv.Itoa(d) + "厘米"
	fmt.Println(aa + "+" + bb + "=")
}

func changdu() {
	mizhuanhua()
	changduyunsuan()
}

// ==============================================================================================================================================================================
func main() {
	fmt.Println("========加减========")
	for i := 0; i < 5; i++ {
		jia()
		jian()
	}
	fmt.Println("========乘法========")
	for i := 0; i < 5; i++ {
		chen()
	}
	fmt.Println("========除法========")
	for i := 0; i < 5; i++ {
		chu()
	}
	fmt.Println("=======元角分=======")
	yuanjiaofen()
	fmt.Println("=======时分秒=======")
	shifenmiao()
	fmt.Println("======长度单位======")
	changdu()
}
