package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
)

func main() {
	bs, err := ioutil.ReadFile("./data.txt")
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	buf := bytes.NewBuffer(bs)
	scanner := bufio.NewScanner(buf)

	count := 0
	elves := make(map[int]int)
	for scanner.Scan() {

		text := scanner.Text()

		if text != "" {
			v, err := strconv.Atoi(text)
			if err != nil {
				log.Printf("Unable to parse: %s", text)
			}
			elves[count] += v
		} else {
			count++
		}
	}

	out := make([]int, 0)
	for _, v := range elves {
		out = append(out, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(out)))
	log.Printf("%v", out[:3])

	sum := 0
	for _, v := range out[:3] {
		sum += v
	}

	log.Printf("result is: %v", sum)
}
