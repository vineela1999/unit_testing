
package main
import (
	
	"testing"
)

type mock_structure struct{}

var mock_function func (string)string

func (m mock_structure)symbol_checker_method(input_string string)string{
	return mock_function(input_string)
}


func TestCategoryFunction(t *testing.T){
	
	sample_cases:= []struct{
		inp_string string
		append_string string
		exp_output string
	}{
		{"vineeLa",""," vineeLa"},
		{"vineela","name","name vineela"},
		{"9837953","number","number 9837953"},
		{"Vineela132","ID","ID Vineela132"},
		
	}
	check_variable=mock_structure{}
	for _,case_value:= range sample_cases{
		
		mock_function = func (input_string string)string{
			return case_value.append_string
		}
	
		output,_:= category_function(case_value.inp_string)
		if(output!=case_value.exp_output){
			t.Errorf("Expected '%s' got '%s'",case_value.exp_output,output)
		}

	}
}

func TestCategoryFunctionNil(t *testing.T){	
	check_variable=mock_structure{}
	mock_function = func (input_string string)string{
	return "name"
	}
	_,error_type:= category_function("")
	if(error_type==nil){
		t.Errorf("expected 'string is not supposed to be empty' got  %s",error_type)
	}
}



