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

	num := 0
	elves := make(map[int]int)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			num++
			continue
		}

		v, err := strconv.Atoi(text)
		if err != nil {
			log.Printf("Unable to parse calories value: %s", text)
		}
		elves[num] += v
	}

	out := make([]int, 0)
	for _, v := range elves {
		out = append(out, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(out)))
	log.Printf("Top 3 elves: %v", out[:3])

	top3 := out[:3]
	sum := 0
	for _, v := range top3 {
		sum += v
	}
	log.Printf("result is: %v", sum)
}
