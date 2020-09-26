
//sum of 2 integers is output
package main

import "fmt"

func sum_of_two_integers(first_integer,second_integer int)int64{
return int64(first_integer+second_integer)
}

func main(){
var first_integer,second_integer int
fmt.Println("Give 2 integers in integer range")
fmt.Scanf("%d",&first_integer)
fmt.Scanf("%d",&second_integer)

fmt.Printf("Sum of given 2 numbers is %d",sum_of_two_integers(first_integer,second_integer))
}

