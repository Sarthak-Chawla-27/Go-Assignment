package model

type Student struct {
	ID    int
	Name  string
	Age   int
	Email string
}

var Students = []Student{
	{ID: 1, Name: "ABC", Age: 20, Email: "abc@gmail.com"},
	{ID: 2, Name: "DEF", Age: 21, Email: "def@gmail.com"},
	{ID: 3, Name: "GHI", Age: 22, Email: "ghi@gmail.com"},
}
