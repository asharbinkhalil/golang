package main

import (
	"fmt"
	"crypto/sha256"
)

// Define the Student struct, which represents a student's information.
type Student struct {
	rollnumber int
	name       string
	address    string
}

// NewStudent is a constructor function that creates a new Student object and returns a pointer to it.
func NewStudent(rollno int, name string, address string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}

// String method for Student to format its output.
func (s *Student) String() string {
	return fmt.Sprintf("student rollno %d\nstudent name %s\nstudent address %s\n", s.rollnumber, s.name, s.address)
}

// Define the StudentList struct, which holds a list of Student objects.
type StudentList struct {
	list []*Student
}

// CreateStudent is a method of StudentList that creates a new Student object, adds it to the list, and returns a pointer to the new Student.
func (ls *StudentList) CreateStudent(rollno int, name string, address string) *Student {
	st := NewStudent(rollno, name, address)
	ls.list = append(ls.list, st)
	return st
}
func calculateHash(student *Student)string{
	data:= fmt.Sprintf("%d%s%s",student.rollnumber,student.name,student.address)
	hash:=sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x",hash)
}

func main() {
	// Create a new StudentList object.
	student := new(StudentList)

	// Create two Student objects and add them to the list using the CreateStudent method.
	student.CreateStudent(24, "Asim", "AAAAAA")
	student.CreateStudent(25, "Naveed", "BBBBBB")

	// Print the list of students with a formatted output.
		for _, student:= range student.list{
			hash:=calculateHash(student)
			fmt.Printf("Student  Hash %s\n",hash)
		}
	
}
