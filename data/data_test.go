package data

import (
	"testing"
)

func TestValidation(t *testing.T){
	p:=&Product{
		Name:"Kate Winslet",
		CreatedOn:"2021-10-18",
		Email:"kate@gmail.com",
		Price:12.5,
	}
	err:=p.Validator()

	if err!=nil{
		t.Fatal(err)
	}
}