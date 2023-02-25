package main

import (
	"bufio"
	"fmt"
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
	var seatLen int = -1
	var reqLen int = -1
	var seats []string

	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}

		if seatLen == -1 {
			seatLen, err = strconv.Atoi(data[:len(data)-1])
			if err != nil {
				log.Fatal(err)
			}
			seats = make([]string, seatLen)
			continue
		}

		if seatLen > 0 {
			seats[cap(seats)-seatLen] = data
			seatLen--
			continue
		}

		if reqLen == -1 {
			reqLen, err = strconv.Atoi(data[:len(data)-1])
			if err != nil {
				log.Fatal(err)
			}
			continue
		}

		if reqLen > 0 {
			ok := false
			buffer := strings.Split(data[:len(data)-1], " ")
			num, err := strconv.Atoi(buffer[0])
			if err != nil {
				log.Fatal(err)
			}

			if num > 3 {
				log.Fatal("wrong num")
			}

			var isLeft bool = false

			if buffer[1] == "left" {
				isLeft = true
			}

			if !isLeft && (buffer[1] != "right") {
				log.Fatal("wrong side")
			}

			var isAisle bool = false

			if buffer[2] == "aisle" {
				isAisle = true
			}

			if !isAisle && (buffer[2] != "window") {
				log.Fatal("wrong position")
			}

			for row := 0; row < cap(seats); row++ {

				if isLeft {
					if isAisle {
						// left aisle
						buffer := num
						var result string
						for site := 0; site < num; site++ {
							if buffer == 0 {
								break
							}

							if seats[row][2-site] == byte('.') {
								buffer--
								result = (" " + fmt.Sprint(row+1) + string(int('C')-site)) + result
							} else {
								buffer = num
							}
						}

						if buffer == 0 {
							writer.WriteString("Passengers can take seats:" + result + "\n")
							writer.Flush()
							for site := 0; site < num; site++ {
								buffer := []rune(seats[row])
								buffer[2-site] = 'X'
								seats[row] = string(buffer)
							}
							for _, str := range seats {
								writer.WriteString(str)
							}
							writer.Flush()
							for site := 0; site < num; site++ {
								buffer := []rune(seats[row])
								buffer[2-site] = '#'
								seats[row] = string(buffer)
							}
							ok = true
							break
						}
					} else {
						// left window
						buffer := num
						var result string
						for site := 0; site < num; site++ {
							if buffer == 0 {
								break
							}
							if seats[row][site] == byte('.') {
								buffer--
								result += (" " + fmt.Sprint(row+1) + string(int('A')+site))
							} else {
								buffer = num
							}
						}

						if buffer == 0 {
							writer.WriteString("Passengers can take seats:" + result + "\n")
							writer.Flush()
							for site := 0; site < num; site++ {
								buffer := []rune(seats[row])
								buffer[site] = 'X'
								seats[row] = string(buffer)
							}
							for _, str := range seats {
								writer.WriteString(str)
							}
							writer.Flush()
							for site := 0; site < num; site++ {
								buffer := []rune(seats[row])
								buffer[site] = '#'
								seats[row] = string(buffer)
							}
							ok = true
							break
						}
					}
				} else {
					if isAisle {
						// right aisle
						buffer := num
						var result string
						for site := 0; site < num; site++ {
							if buffer == 0 {
								break
							}
							if seats[row][4+site] == byte('.') {
								buffer--
								result += (" " + fmt.Sprint(row+1) + string(int('D')+site))
							} else {
								buffer = num
							}
						}

						if buffer == 0 {
							writer.WriteString("Passengers can take seats:" + result + "\n")
							writer.Flush()
							for site := 0; site < num; site++ {
								buffer := []rune(seats[row])
								buffer[4+site] = 'X'
								seats[row] = string(buffer)
							}
							for _, str := range seats {
								writer.WriteString(str)
							}
							writer.Flush()
							for site := 0; site < num; site++ {
								buffer := []rune(seats[row])
								buffer[4+site] = '#'
								seats[row] = string(buffer)
							}
							ok = true
							break
						}
					} else {
						// right window
						buffer := num
						var result string
						for site := 0; site < num; site++ {
							if buffer == 0 {
								break
							}

							if seats[row][6-site] == byte('.') {
								buffer--
								result = (" " + fmt.Sprint(row+1) + string(int('F')-site)) + result
							} else {
								buffer = num
							}
						}

						if buffer == 0 {
							writer.WriteString("Passengers can take seats:" + result + "\n")
							writer.Flush()
							for site := 0; site < num; site++ {
								buffer := []rune(seats[row])
								buffer[6-site] = 'X'
								seats[row] = string(buffer)
							}
							for _, str := range seats {
								writer.WriteString(str)
							}
							writer.Flush()
							for site := 0; site < num; site++ {
								buffer := []rune(seats[row])
								buffer[6-site] = '#'
								seats[row] = string(buffer)
							}
							ok = true
							break
						}
					}
				}
			}
			if !ok {
				writer.WriteString("Cannot fulfill passengers requirements\n")
				writer.Flush()
			}
			reqLen--
			continue
		}
	}
}
