package main

import (
	"fmt"
	"strconv"
)

func main() {
	//runeFunc("heelo 北京")
	typeCast("1234")
}

func runeFunc(str2 string) {
	str3 := []rune(str2)
	for i := 0; i < len(str3); i++ {
		fmt.Printf("%c", str3[i])
	}
}

func typeCast(str string) {
	n, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("typeCastError", err)
	} else {
		fmt.Println("typeCastSuccess", n)
	}
}
