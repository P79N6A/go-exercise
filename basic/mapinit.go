package basic

import "fmt"

func InitSlice()  {
	slice1:=[]int{1,2,3,4,5,6,7,8,9,0}
	slice2:=make([]*int,len(slice1))
	for index,item:=range slice1{
		slice2[index]=&item
		fmt.Println("size=", len(slice2),"index=",index)
	}
}
