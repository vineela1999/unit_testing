package main
import "fmt"


type student struct{
	student_name string
	ID string
	marks int
	result string
}


type studentInf interface {
	student_structure_creation(string,string,int) student
}

type student_structure struct{}

func (s student_structure) student_structure_creation(name string,id string,mark int) student {
	
	return structure_function(name,id,mark)
}

var student_var studentInf

func init() {
	student_var = student_structure{}
}


func structure_function(name string,id string,mark int)student{		//student details are stored as object
	var student_object student
	student_object.student_name=name
	student_object.ID=id
	student_object.marks=mark

	return student_object
}
	
func result_generator(name string,id string,mark int)(string,error){   //gives grade for the marks 
	
	if(mark<0 || mark>100){
		return "",fmt.Errorf("Marks should be in range of 1-100")	
	}
	student_object:= student_var.student_structure_creation(name,id,mark) 
	
	if(mark>=0 && mark<30){
		student_object.result="fail"	
	}

	if(mark>=30 && mark<50){
		student_object.result="grade C"	
	}

	if(mark>=50 && mark<80){
		student_object.result="grade B"	
	}
	
	if(mark>=80 && mark<=100){
		student_object.result="grade A"	
	}

	return student_object.result,nil

}



func main(){
var name,id string
var mark int

fmt.Println("Enter name of student")
fmt.Scanf("%s",&name)
fmt.Println("Enter ID of student")
fmt.Scanf("%s",&id)
fmt.Println("Enter marks of student")
fmt.Scanf("%d",&mark)

result,error_type:=result_generator(name,id,mark)

if(error_type!=nil){
	fmt.Println(error_type)
}

if(error_type==nil){
	fmt.Printf("%s result is: %s \n",name,result) 
}


}
