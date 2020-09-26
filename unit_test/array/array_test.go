package main
import "testing"
import "fmt"

type mock_interface interface{
	sort_arrary_method([]int)
}

type mock_structure struct{}

var mock_function func([]int)

func (m mock_structure)sort_arrary_method(input_slice []int){
	mock_function(input_slice)
}


func Test_array_max_frequency(t *testing.T){
	sortFuncInf_var=mock_structure{}
	
	
	mock_function= func (input_slice []int){
			sorted_output:=[]int{1,1,1,3,4,4,5}
			input_slice=sorted_output
			fmt.Println("in test func",input_slice)
		}
		slice:=[]int{4,5,1,4,1,1,3}
	frequency:=array_max_frequency(slice)
	if(frequency!=3){
		t.Errorf("Expected 3 got %d",frequency)
	}

}
