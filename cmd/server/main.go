package main

import (
	"fmt"
	"vivaop/internal/util"
)

func main() {
	fmt.Println(util.RandomString(32))
	fmt.Println(util.RandomEmail())
	fmt.Println(util.HashPassword("secret"))
}
