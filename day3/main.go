package main

import (
	"bufio"
	"log"
	"os"
)

func chal1() error {
	f, err := os.Open("./input1.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		line := sc.Text()
		log.Println(line)
	}

	return sc.Err()
}

func main() {
	if err := chal1(); err != nil {
		log.Fatalln(err)
	}
}
