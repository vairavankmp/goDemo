package main

import (
	"fmt"
)

type dividedbyzero struct{
	texterror string
}

func(d dividedbyzero)Error() string{
	return d.texterror
}
func divideerror()error{
	return dividedbyzero {"divide by zero using structure"}
}
 

func getdivision(num1, num2 int) (int,error) {
	if num2==0{
		return -1,divideerror()
	}
	return num1 / num2,nil

}
func main() {

	c ,err := getdivision(10, 0)
	if err!=nil{
		fmt.Println(err.Error())
	}else{
	fmt.Println(c)
	}
}