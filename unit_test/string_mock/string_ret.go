
package main

import "fmt"
//import "strings"


type symbol_Checker_interface interface {
	symbol_checker_method(string) string
}

type symbol_Checker_struct struct{}

func (a symbol_Checker_struct) symbol_checker_method(input_string string) string {
	
	return symbol_checker_function(input_string)
}

var check_variable symbol_Checker_interface

func init() {
	check_variable = symbol_Checker_struct{}
}


func symbol_checker_function(input_string string)string{

	length_of_string:= len(input_string)
	
	smalls_of_string,capitals_of_string,digits_of_string:=0,0,0
	for _,symbol:= range input_string{
		if((symbol>='a' && symbol<='z')){
			smalls_of_string++
		}
		if((symbol>='A' && symbol<='Z')){
			capitals_of_string++
		}
		if((symbol>='0' && symbol<='9')){
			digits_of_string++
		}
		
	}
	
	

	if(digits_of_string<length_of_string && digits_of_string>0){		//If combination of numbers and letters ID id appended
		return "ID"
	}

	if(digits_of_string==length_of_string){		//If combination of numbers number id appended
		return "number"
	}
	if(smalls_of_string==length_of_string){		//If combination of only small letters name appended
			return "name"
		}
	
return ""						// if others nothingis appended
	
}



func category_function(input_string string) (string,error){
//fmt.Println("came to my function")
value:= check_variable.symbol_checker_method(input_string)

if(input_string==""){
	return input_string,fmt.Errorf("string is not supposed to be empty")
}

appended_string:=value+" "+input_string
	return appended_string,nil


}

func main(){
	var input_string string
	fmt.Println("Enter a string an non empty string :")
	fmt.Scanf("%s",&input_string)
	output_string,error_type:= category_function(input_string)
	fmt.Println(error_type)
	if(error_type==nil){
		fmt.Println(output_string)
		
	}
	if(error_type!=nil){
		
		fmt.Println(error_type)
	
	}
}











	

