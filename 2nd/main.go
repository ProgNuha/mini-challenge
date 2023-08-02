package main

import (
    "fmt"
    "os"
	"2nd/helpers"
	"strconv"
)

func main() {
	person := helpers.Participant{}
	argsRaw := os.Args
	args := argsRaw[1:]
	length := len(args)
	
	if (length == 0) {
		fmt.Printf("Input id peserta") 
	} else if(length == 1){
		idPerson, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		  }
		person.GetParticipant(idPerson)
	}
}
