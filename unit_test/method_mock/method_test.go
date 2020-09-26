package main

import (
    "testing"
)


type MockTest struct {
    Elem int
}

// GetVolume returns the volume of a cube
func (config *MockTest) area() int {
    return config.Elem * config.Elem
}

// GetSurfaceArea returns the surface area of a cube
func (config *MockTest) perimeter() int {
    return  4* config.Elem 
}


func TestGetVolumeAndArea(t *testing.T){
	sample_cases:= []struct{

	 input MockTest
	 exp_area int
	 exp_perimeter int	 
	}{
		{MockTest{2},4,8},
 		{MockTest{1},1,4},
		{MockTest{11},121,44},
	
	}

	for _,case_value:= range sample_cases{
		 area,perimeter := GetareaAndperimeter(&(case_value.input))
		
		if (area!= case_value.exp_area) {
                t.Errorf("Expected volume to be equal to %d but was equal to %d ",
                    area, case_value.exp_area)
                
            }

            if (perimeter!= case_value.exp_perimeter) {
                t.Errorf("Expected area to be equal to %d but was equal to %d ",
                    perimeter, case_value.exp_perimeter)
                
	    }	

}


}



