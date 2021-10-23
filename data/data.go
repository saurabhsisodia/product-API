package data

import (
	"strconv"
	"errors"
	"github.com/go-playground/validator"
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
)

type Product struct{
	Id string 			`json:"id"`
	Name string			`json:"name" validate:"required"`
	CreatedOn string 	`json:"createdon" validate:"required"`
	Email string     	`json:"email" validate:"required,email"`
	Price float32    	`json:"price" validate:"required,gt=0"`
}

var productList = []Product{
	{Id:"1",Name:"Coffee",CreatedOn:"2021-10-17",Email:"coffe@gmail.com",Price:20.00},
	{Id:"2",Name:"Tea",CreatedOn:"2021-10-17",Email:"tea@gmail.com",Price:120.00},
	{Id:"3",Name:"Beer",CreatedOn:"2021-10-17",Email:"beer@gmail.com",Price:120.00},
}

func (p *Product) Validator()error{
	validate:=validator.New()
	return validate.Struct(p)
}


func GetProduct(id string)*Product{
	for _,u:=range productList{
		if u.Id==id{
			return &u
		}
	}
	return &Product{}
}
func GetProducts()*[]Product{
	return &productList
}

func AddProduct(prod *Product){
	prod.Id=strconv.Itoa(GetId())
	productList=append(productList,*prod)
}
func GetId()int{
	v,_:=strconv.Atoi(productList[len(productList)-1].Id)
	return v+1
}
func UpdateProduct(d Product)error{
	for i:=0;i<len(productList);i++{
		if productList[i].Id==d.Id{
			productList[i]=d
			return nil
		}
	}
	return errors.New("Product does not exist")
}
func DeleteProduct(id string)error{
	ind:=-1
	for i:=0;i<len(productList);i++{
		if productList[i].Id==id{
			ind=i
			break
		}
	}
	if ind==-1{
		return errors.New("Product does not exist")
	}
	productList=append(productList[:ind],productList[ind+1:]...)
	return nil
}

func ValidationMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(

		func (rw http.ResponseWriter,r *http.Request){

			prod:=&Product{}
			//err:=json.NewDecoder(r.Body).Decode(prod)
			body,_:=ioutil.ReadAll(r.Body)
			err:=json.Unmarshal(body,prod)
			if err!=nil{
				http.Error(rw,"error while decoding",http.StatusBadRequest)
				return
			}

			err = prod.Validator()
			if err!=nil{
				http.Error(rw,fmt.Sprintf("%v",err),http.StatusBadRequest)
				return
			}
			r.Body=ioutil.NopCloser(bytes.NewBuffer(body))  // set the request body again to read by next 
															// handler or middleware
			next.ServeHTTP(rw,r)
		},
	)
}