package entities

type Employee struct {
	Id    int32
	Name  string
	Age int32
}

func NewEmployee(name string,age int32) *Employee {
	return &Employee{Id: 1, Name: name, Age: age}
}