package main

import (
	"bufio"
	"log"
	"os"
)

func Intersect(line string, line string) []string {


}

func FindOnlyItem(line string) string {
    w := int(len(line) / 2)
    first := line[:w]
    second := line[w:]
    for _, c := first {
        if c := 
    }
}

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
