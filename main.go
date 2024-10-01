package main

import (
	"bufio"
	// "bytes"
	"flag"
	"fmt"
	"os"
	"strings"
)


	func word_count(data string) int {

		sum := 0
		words := strings.Fields(data)

		for _, word := range words {
			fmt.Println(word)
			sum++
		}
		// fmt.Println(sum)
		return sum
	
	}

	func char_count(data string) int {
		sum := 0

		for _, char := range data {
			sum++
			fmt.Println(char)
		}
		return sum
	}

	func line_count(data string) int {
		number_of_lines := 0
		lines := strings.Split(data, "\n")

		for _, line := range lines {
			fmt.Println(line)
			number_of_lines ++
		}
		return number_of_lines
	}

	func main(){
		reader := bufio.NewReader(os.Stdin)

		words := flag.Bool("w", false, "counts words")
		bytes := flag.Bool("c", false, "counts bytes")
		lines := flag.Bool("l", false, "counts lines")
		characters := flag.Bool("m", false, "counts characters")
		defaultt := flag.Bool("", false, "counts characters")



		flag.Parse()

		fmt.Println("Hey Please enter your file")
		filename, err := reader.ReadString('\n')
		// fmt.Println(filename)


		if err != nil {
			fmt.Println(err)
		}

		filename = strings.TrimSpace(filename)

		content, err := os.ReadFile(filename)
		// fmt.Println(content)


		data := string(content)

		// fmt.Println(data)

		// fmt.Println("file is " + filename)


		// * is a pointer and now we are derefencing
		if *defaultt == true {
			fmt.Println(word_count(data), len(content), line_count(data), char_count(data))
		}
		if *words == true {
			// fmt.Println("we made it to word count")
			word_count := word_count(data)
			fmt.Println(word_count)
			// count all the words 
		} 
		if *bytes == true {

			byte_sum := len(content)
			fmt.Println(byte_sum)
		
		} 
		if *lines == true {
			line_count := line_count(data)
			fmt.Println(line_count)
		} 
		
		if *characters == true {
			char_count := char_count(data)
			fmt.Println(char_count)
		} else {
			fmt.Println(word_count(data), len(content), line_count(data), char_count(data), filename)

		}



		
	}