package main

import (
	"fmt"
	"crypto/sha256"
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
	return fmt.Sprintf("student rollno %d\nstudent name %s\nstudent address %s\n", s.rollnumber, s.name, s.address)
}

type StudentList struct {
	list []*Student
}

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
	student := new(StudentList)

	student.CreateStudent(24, "Asim", "AAAAAA")
	student.CreateStudent(25, "Naveed", "BBBBBB")

		for _, student:= range student.list{
			hash:=calculateHash(student)
			fmt.Printf(student.name)
			fmt.Printf(" %d ",student.rollnumber)
			fmt.Printf(student.address)
			fmt.Printf("\nHash %s\n",hash)
		}
	
}
