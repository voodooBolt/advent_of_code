package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	_ "reflect"
	"strconv"
	_ "slices"
	"sort"
	"math"
)

//var input_file string = "./test.txt"
var input_file string = "./puzzle_input.txt"


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
	slice1 := make([]int, 0)
	slice2 := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		
		// grab each number
		num1_s := strings.Split(line, "   ")[0]
		num2_s := strings.Split(line, "   ")[1]
		
		// convert to int
		num1, err := strconv.Atoi(num1_s)
		check(err)
		num2, err := strconv.Atoi(num2_s)
		check(err)

		//fmt.Println(reflect.TypeOf(num1)) // prints the variable type
		//fmt.Printf("%d , %d\n",num1,num2)

		// append to slice
		slice1 = append(slice1, num1)
		slice2 = append(slice2, num2)
	}

	//fmt.Println(slice1)
	//fmt.Println(slice2)

	// sort slices
	sort.Ints(slice1)
	sort.Ints(slice2)
	
	// calculate distance on each slice and add to sum
	total_distance := 0.0
	for i, _ := range slice1{
		distance := math.Abs(float64(slice1[i] - slice2[i]))
		total_distance += distance
		//fmt.Println(distance)
	}
	fmt.Printf("%f\n", total_distance)


	fmt.Println("done")
}

