package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		counts[text]++
		count := counts[text]
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, text)
		}
	}

}
