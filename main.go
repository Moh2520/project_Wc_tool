package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)


func word_count(word, string) int {




}

func main(){
	// read in file 
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hey Please enter your file")
	filename, _ := reader.ReadString('\n')

	fmt.Println("file is" + filename)

	words := flag.Bool("w", false, "counts words")
	// * is a pointer and now we are derefencing
	if *words {
		// count all the words 


	}

}