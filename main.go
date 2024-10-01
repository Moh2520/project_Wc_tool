package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)


	func word_count() int {

		sum := 0
		data := "This is a sample sentence"
		words := strings.Fields(data)

		for _, word := range words {
			fmt.Println(word)
			sum++
		}
		fmt.Println(sum)
		return sum
	
	}

func main(){
	read in file 
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hey Please enter your file")
	filename, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	// content, err := os.ReadFile(filename)
	

	// data := string(content)

	fmt.Println("file is" + filename)

	words := flag.Bool("w", false, "counts words")

	flag.Parse()
	// * is a pointer and now we are derefencing
	if *words {
		word_count()
		// count all the words 

	word_count()
	}


}