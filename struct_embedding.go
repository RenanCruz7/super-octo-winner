package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person // Embedding Person struct in Employee
	Job    string
}

func main() {
	e := Employee{
		Person: Person{Name: "Alice", Age: 30},
		Job:    "Engineer",
	}

	fmt.Println("Name:", e.Name)
	fmt.Println("Age:", e.Age)
	fmt.Println("Job:", e.Job)
}
