
package main 

import "testing"


type sample_cases struct{

	first_no int
	second_no int
	expected_output int64
}


func TestSum_of_two_integers(t *testing.T){

sample_cases:= []struct{

	first_no int
	second_no int
	expected_output int64
}{
	{20,30,50},
 	{-561,30,-531},
	{-64,0,-64},
	{-43,-300,-343},
	{21474483647,21474483647,42948967294},
	{-21474483647,-21474483647,-42948967294},
}

	for _,samples_value:=range sample_cases{
		output_sum:= sum_of_two_integers(samples_value.first_no,samples_value.second_no)
		if(output_sum!=samples_value.expected_output){
			t.Errorf("expected %d but got as %d",samples_value.expected_output,output_sum)
		}
	}
}


