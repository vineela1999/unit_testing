
package main

import "fmt"

//creating an interface to implement function small_letter_checker function
type small_letterCheckerInf interface {
	small_letter_checker(string) bool
}

//creating a structure to implement the small_letterCheckerInf interface
type small_letterCheck struct{}

//here this method calls small_letter_check function
func (r small_letterCheck) small_letter_checker(input_string string) bool {
	
	return small_letter_check(input_string)
}

//a variable of type interface is created
var check_var small_letterCheckerInf

func init() {
	check_var = small_letterCheck{}
}

//this functions checks whether given string is total of small letters or not
func small_letter_check(input_string string)bool{

	length_of_string:= len(input_string)
	fmt.Println(length_of_string)
	iterator:=0
	for _,character:= range input_string{
		if(!(character>='a' && character<='z')){
			break
		}
		iterator++
	}
	fmt.Println(iterator)
	if(iterator==length_of_string){
		return true		// if consists of small letters only
	}

	if(iterator!=length_of_string){
		return false		//if any other literal exists
	}
	return true
}



func small_letters_function(input_string string) (string,error){

value:= check_var.small_letter_checker(input_string)

if(value==true){		//if only small letters then hi is appended and returned
	
	input_string= "hi "+ input_string		
	return input_string,nil
}

if(value==false){		//if any other literal exists, error is returned
	
	return input_string,fmt.Errorf("string '%s' is not in format", input_string) 	
}
	return input_string,nil
}

func main(){
	var input_string string
	
	fmt.Scanf("%s",&input_string)
	output_string,func_error:= small_letters_function(input_string)
	
	if(func_error==nil){
		fmt.Println(output_string)
		
	}
	if(func_error!=nil){
		fmt.Println(func_error)
	}
}



