package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/magicalbanana/peggy-hill/golang/peggy"
)

var digitsRegex = regexp.MustCompile(`[0-9]+`)
var removeEmptyLineRegex = regexp.MustCompile(`[\t\r\n]+`)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("no pipe :( you need to pipe it sir")
	} else {
		reader := bufio.NewReader(os.Stdin)
		buf := make([]byte, 0, 4*1024)
		pegs := make([]int, 0)
		for {
			n, err := reader.Read(buf[:cap(buf)])
			buf = buf[:n]
			if n == 0 {
				if err == nil {
					continue
				}
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			// process buf
			if err != nil && err != io.EOF {
				log.Fatal(err)
			}
			// we need to sanitize the lines to remove the extra
			// white space. Not sure why there's an extra white
			// space
			lines := strings.Split(removeEmptyLineRegex.ReplaceAllString(strings.TrimSpace(string(buf)), "\n"), "\n")
			for _, line := range lines {
				if !digitsRegex.Match([]byte(line)) {
					log.Fatal("Can only process numbers, regex: [0-9]+")
				}
				i, err := strconv.Atoi(line)
				if err != nil {
					log.Fatal(err)
				}
				if i == 0 {
					return
				}
				pegs = append(pegs, i)
			}
		}

		log.Println("Given pegs:", pegs)
		answer := peggy.Calculate(pegs)
		log.Println("Gears:", answer)
	}
}
