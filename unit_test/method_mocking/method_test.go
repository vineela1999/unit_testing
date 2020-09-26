package main

import (
    "testing"
)


type MockTest struct {
    length int
	breadth int
}

var mock_func_area func()int
var mock_func_perimeter func()int


func (config *MockTest) area() int {
    return mock_func_area()
}


func (config *MockTest) perimeter() int {
    return  mock_func_perimeter() 
}



func TestGetVolumeAndArea(t *testing.T){

	sample_cases:= []struct{

	 input MockTest
	 exp_area int
	 exp_perimeter int
	exp_sum int
	 exp_product int	
		 
	}{
		{MockTest{3,4},12,14,26,168,},
 		{MockTest{1,1},1,4,5,4},
		{MockTest{6,1},6,14,20,84},
	
	}



for _,case_value:= range sample_cases{	
	mock_func_area =func()int{
		return case_value.exp_area
	}

	mock_func_perimeter =func()int{
		return case_value.exp_perimeter
	}
	sum,product := GetareaAndperimeter(&case_value.input)
	
		if (sum!=case_value.exp_sum) {
                t.Errorf("Expected sum to be equal to %d but was equal to %d ",case_value.exp_sum,sum)
                }

            if (product!=case_value.exp_product) {
                t.Errorf("Expected product to be equal to %d but was equal to %d ",case_value.exp_product,product)
                }
}	

}


