package main

import "fmt"

func main() {
	println("Start!!!")
	str := "Hi, good morning!"

	passByReference(&str)

	passByValue(str)
}

func passByReference(addStr *string) {
	myStringAddress := fmt.Sprintf("%p", addStr)
	println("get address -> |" + myStringAddress + "| and acessing the address we has |\"" + *addStr + "\"|")
}

func passByValue(valueStr string) {
	myStringAddress := fmt.Sprintf("%p", &valueStr)
	println("get value -> |" + valueStr + "| and getting the address we has |\"" + myStringAddress + "\"|")
}
