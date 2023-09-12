package main
import "fmt"

type employee struct
{
	name string
	salary int
	position string
}
type company struct
{
	name string
	employees []employee

}

func main(){
	emp1 := employee{"Amir", 80000, "Full-Stack Developer"}
	emp2 := employee{"Ali", 60000, "Front-End Developer"}
	emp3 := employee{"Ahmed", 50000, "Back-End Developer"}


	emplys:=[]employee{emp1,emp2,emp3}
	comp:=company{"Google",emplys}
	fmt.Println(comp)
}