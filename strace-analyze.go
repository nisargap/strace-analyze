// Name: Nisarga Patel
// Class: CMSC 491

package main

import "fmt"
import "os"
import "bufio"
import "strings"

var commandMap = make(map[string]int)

func readLine(path string) {

	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	var currentcommand = ""
	var linecount = 0
	for scanner.Scan() {
		linecount++
		currentcommand = strings.Split(scanner.Text(), "(")[0]
		if _, ok := commandMap[currentcommand]; ok {
			commandMap[currentcommand] += 1
		} else {
			commandMap[currentcommand] = 0
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("Number of lines %d\n", linecount)
	fmt.Println("Command, Frequency")
	for key, value := range commandMap {

		fmt.Printf("%s, %d\n", key, value)

	}
	fmt.Println("---------------------------------")

}
func main() {

	if len(os.Args) == 2 {
		readLine(os.Args[1])
	} else {
		fmt.Println("Please enter absolute file path as comandline argument!")
	}
}
