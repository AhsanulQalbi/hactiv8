package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	name    string
	address string
	job     string
	reason  string
}

func main() {
	students := []Student{
		{name: "Abraham Purnomo", address: "Tangerang", job: "Backend-Engineer", reason: "ingin menambah pengalaman"},
		{name: "Ammar Farghani", address: "Malang", job: "Backend-Engineer", reason: "ingin menambah pengalaman"},
		{name: "Muhammad Zulfa Azhari", address: "Jogja", job: "Backend-Engineer", reason: "ingin menambah pengalaman"},
		{name: "Raka Adli Pramudita", address: "Jakarta Timut", job: "Backend-Engineer", reason: "ingin menambah pengalaman"},
	}
	argsWithProg := os.Args[1]
	i, _ := strconv.Atoi(argsWithProg)
	printStudent(students[i-1])
}

func printStudent(student Student) {
	fmt.Printf("%+v", student)
}
