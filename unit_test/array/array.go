package main

import "fmt"


type sortFuncInf interface{
	sort_arrary_method([]int)
} 

type array_structure_type struct{}

var sortFuncInf_var sortFuncInf

func (a array_structure_type)sort_arrary_method(input_slice []int){
	sort_arrary_function(input_slice)
} 


func init(){
	sortFuncInf_var=array_structure_type{}
}


func sort_arrary_function(input_slice []int){
	for position:=0;position<(len(input_slice)-1);position++{
		for iterator:=position+1;iterator<len(input_slice);iterator++{
			
			if(input_slice[position]>input_slice[iterator]){
				variable:=input_slice[iterator]
				input_slice[iterator]=input_slice[position]
				input_slice[position]=variable
			}
			
		}
		
	}
}


func array_max_frequency(input_slice []int)int{
	
	sortFuncInf_var.sort_arrary_method(input_slice)
	fmt.Println(input_slice)
	
	var present_max int=1;
	iterator:=1
	for position:=0;position<(len(input_slice)-1);position++{
		if(input_slice[position]==input_slice[position+1]){
			iterator++;
		}
		if(input_slice[position]!=input_slice[position+1]){
			if(iterator>present_max){
				present_max=iterator
				iterator=1			
			}		
		}	

	}
	if(iterator>present_max){
		present_max=iterator
	}
	fmt.Println(present_max)
	return present_max
}


func main(){
	var number_of_entries int
	var input_array [100]int
	fmt.Println("Enter number of entries: ")
	fmt.Scanf("%d",&number_of_entries)
	fmt.Println("Enter elements of input_array: ")
	for position:=0;position<number_of_entries;position++{
		fmt.Scanf("%d",&input_array[position])
	}

	input_slice:= input_array[:number_of_entries]
	
	max_frequency:=array_max_frequency(input_slice)
	fmt.Println(max_frequency)
}
