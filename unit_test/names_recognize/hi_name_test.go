package main

import "testing"
//import "fmt"


var mock_function func (string)bool

type mock_structure struct{}

func (m mock_structure)small_letter_checker(input_string string) bool{
	return mock_function(input_string)
}

//To test whether appending hi to the only small letters string
func TestMy_function_perfect(t *testing.T){


sample_cases:= []struct{

	input_string,expected_output string
	}{
		{"vineela","hi vineela"},
 		{"dgh","hi dgh"},
		{"p vineela","hi p vineela"},
	
}	

	check_var= mock_structure{}	
	mock_function = func(input_string string) bool {
		return true
	}

	for _,sample_value:= range sample_cases{
	
	output_string,error_type:=small_letters_function(sample_value.input_string)

		if(output_string!=sample_value.expected_output){
			t.Errorf("wrong output expected %s got %s",sample_value.expected_output,output_string)	
		}
	
		if(error_type!=nil){
			t.Error("error should be a nil")	
		}

	}
}

//To test whether showing error to the only small letters string
func TestMy_function_imperfect(t *testing.T){


	check_var= mock_structure{}	
	mock_function = func(input_string string) bool {
		return false
	}
		
	_,error_type:=small_letters_function("vineela123")
	
	if(error_type==nil){
		t.Errorf("error shouldnt be a nil expecting an error %s",error_type)	
	}
}


