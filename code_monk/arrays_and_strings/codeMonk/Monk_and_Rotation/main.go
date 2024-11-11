package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numTestCase int
	var n, k int
	var line string
	_, err := fmt.Scanf("%d", &numTestCase)
	if err != nil {
		return
	}
	reader := bufio.NewReader(os.Stdin)

	for numTestCase > 0 {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		n, _ = strconv.Atoi(parts[0])
		k, _ = strconv.Atoi(parts[1])

		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		arr := strings.Split(line, " ")

		index := n - (k % n)

		for i := index; i < len(arr); i++ {
			fmt.Print(arr[i], " ")
		}

		for j := 0; j < index; j++ {
			fmt.Print(arr[j], " ")
		}

		fmt.Println()
		numTestCase--
	}
}
