//include main package to ensure that program will start to main
package main

//import the library fmt use for Print function
import "fmt"

//main function
func main() {
//create an array size 5
    arr := [5]int{1,2,3,4,5}
//print the element which are exis in arr array variable
    fmt.Println("array element \t:\t", arr)

//create slice_1 using arr array start from index 2 to last index
    slice_1 := arr[2:]
//print the slice_1 element
    fmt.Println("slice_1 element is \t:\t", slice_1)
    
//create slice_2 using arr array start from index 2 to index 3
    slice_2 := arr[2:4]
//print the slice_2 element
    fmt.Println("slice_2 element is \t:\t", slice_2)
}