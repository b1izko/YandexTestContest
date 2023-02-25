package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	out, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	reader := bufio.NewReader(in)
	writer := bufio.NewWriter(out)

	var numbers []int
	var n int = -1
	var k int = -1

	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}

		if n == -1 {
			buffer := strings.Split(data[:len(data)-1], " ")
			n, err = strconv.Atoi(buffer[0])
			if err != nil {
				log.Fatal(err)
			}
			numbers = make([]int, n)

			k, err = strconv.Atoi(buffer[1])
			if err != nil {
				log.Fatal(err)
			}
			continue
		}

		buffer := strings.Split(data[:len(data)-1], " ")
		for i, num := range buffer {
			numbers[i], err = strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}
