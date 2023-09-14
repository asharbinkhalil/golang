package main
import "fmt"


type Person struct
{
    name string
    age int
    weight float32
    height float32
}

func paasingarray() {
    array1:= []string{"ashar", "ali", "khan"}
    array2:= array1
    array3:= &array1

    fmt.Printf("array1: %v\n", array1)
    fmt.Printf("array2: %v\n", array2)
    array1[0]="ahmed"
    fmt.Printf("array1: %v\n", array1)
    fmt.Printf("array2: %v\n", array2)
    fmt.Printf("array3: %v\n", *array3)
}


func arrays() {
    var arr1 = []int{1,2,3}
    arr2 := [5]int{4,5,6,7,8}
    var arr3 = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
    fmt.Println(arr1)
    fmt.Println(arr2)
    fmt.Println(arr3)
    fmt.Println(arr1[0])
    fmt.Println(arr2[2])
    fmt.Println(len(arr2))
  }

func passstructtoarray_task1(obj Person) {
    fmt.Println("Name: ",obj.name)
    fmt.Println("Age: ",obj.age)
    fmt.Println("Weight",obj.weight)
    fmt.Println("Height: ",obj.height)
}

func main() {
    //items := []int{1,2,3,4,5,6,7,8,9,10}
    //fmt.Println(items)

    //name:= "ashar"
    //weight:= 70.5

    //fmt.Printf("my name is %s and my weight is %f\n", name, weight)
    //for i:=0; i<10; i++ {
    //    fmt.Println(items[i])
    //  fmt.Println(" ")
    //}


    //arrays()
    //paasingarray()
    //var obj Person
    //obj.name = "ashar"
    //obj.age = 25
    //obj.weight = 70.5
    //obj.height = 5.8
    //fmt.Println(obj)

    var obj2 = Person{"ali", 25, 70.5, 5.8}
    //fmt.Println(obj2)
    passstructtoarray_task1(obj2)

}
