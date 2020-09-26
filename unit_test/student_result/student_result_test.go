
package main

import "testing"

type mock_structure struct{}

var mock_function func(name string,id string,mark int)student

func (mock_var mock_structure)student_structure_creation(name string,id string,mark int)student{
	return mock_function(name,id,mark)
}


func TestStudent_result_function_gradeB(t *testing.T){

sample_cases:= []struct{

	 name string
	id string
	mark int
	exp_structure student
	exp_output string	
		 
	}{
		{"vineela","b160660cs",94,student{student_name:"vineela",ID:"b160660cs",marks:94},"grade A"},

		{"vineela","b160660cs",67,student{student_name:"vineela",ID:"b160660cs",marks:67},"grade B"},
		{"vineela","b160660cs",45,student{student_name:"vineela",ID:"b160660cs",marks:45},"grade C"},
		{"vineela","b160660cs",23,student{student_name:"vineela",ID:"b160660cs",marks:23},"fail"},

						
	}


	student_var=mock_structure{}
	for _,j:= range sample_cases{
		
	mock_function= func(name string,id string,mark int)student{
		return j.exp_structure
	}

		result,error_type:=result_generator(j.name,j.id,j.mark)

	if(error_type!=nil){
		t.Errorf("expected no error but got %s",error_type)
	
	}

	if(result!=j.exp_output){
		t.Errorf("expected %s but got %s",j.exp_output,result)
	}
	

	}

}


func TestStudent_result_function_range(t *testing.T){		//marks if out of bound

	var sample = student{student_name:"vineela",ID:"b160660cs",marks:-6}


	student_var=mock_structure{}
	mock_function= func(name string,id string,mark int)student{
		return sample
	}

	_,error_type:=result_generator("vineela","b160660cs",-6)

	if(error_type==nil){
		t.Errorf("expected  marks out of range but got %s",error_type)
	
	}

	

}




