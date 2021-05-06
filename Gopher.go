package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	args := os.Args
	csvFile, mError := os.Open(args[1])
	inputReader := bufio.NewReader(os.Stdin)
    
	if (mError != nil) {

		fmt.Println(mError)

	}
	score := 0
	total := 0
	reader := csv.NewReader(csvFile)
	data, merror := reader.Read();
	for merror == nil{

		question := data[0]
		answer :=  data[1] + string(rune(byte(13))) + "\n"
        fmt.Printf("Question #%v: %v = ", total, question)
		userAnswer, _ := inputReader.ReadString('\n')
        
		if userAnswer == answer {
			score++
		}
		total++
		data, merror = reader.Read();

	}
	fmt.Printf("You scored: %v/%v!", score, total) 


}