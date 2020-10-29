package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("filecheck/cs1.txt")
	reader := bufio.NewReader(file)
	flag := 0
	m := 0
	all := make(map[string]string)
	downa := make(map[string]string)
	for {

		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		trimline := strings.Trim(line, "\n")
		if trimline == "===" {
			flag = 1
		}
		if flag == 0 {
			all[line] = line
		}
		if flag == 1 {
			downa[line] = line
		}

	}

	for _, val := range all {
		for i := 0; i < len(downa); i++ {
			if val == downa[val] {
				break
			}
			if i == len(downa)-1 {
				fmt.Println(val)
				m++
			}
		}
	}
	fmt.Println(m)
}
