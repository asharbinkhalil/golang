package main

import (
	"fmt"
)

type Student struct {
	rollnumber int
	name       string
	address    string
}

func NewStudent(rollno int, name string, address string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}

func (s *Student) String() string {
	return fmt.Sprintf("Student rollno %d\nStudent name %s\nStudent address %s\n", s.rollnumber, s.name, s.address)
}

type StudentList struct {
	list []*Student
}
func (ls *StudentList) CreateStudent(rollno int, name string, address string) *Student {
	st := NewStudent(rollno, name, address)
	ls.list = append(ls.list, st)
	return st
}

func main() {
	student := new(StudentList)

	student.CreateStudent(24, "Asim", "AAAAAA")
	student.CreateStudent(25, "Naveed", "BBBBBB")
	for i, st := range student.list {
		fmt.Printf("============ List %d ==========\n", i)
		fmt.Print(st)
		fmt.Println("==============================")
	}
}
