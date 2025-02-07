package main

import(
	"fmt"
	"os"
	"bufio"
)

var input_file string = "./test.txt"
// input_file := "./puzzle_input.txt"


func check(e error){
	if e != nil{
		panic(e)
	}
}

func main() {
	file, err := os.Open(input_file)
	check(err)
	
	defer file.Close()

	//new scanner
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	
	//process data into 2 ordered lists
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}


	fmt.Println("done")
}

