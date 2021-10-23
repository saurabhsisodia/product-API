package handlers

import (
	"fmt"
	"log"
	"net/http"
	"context"

	"REST-API/data"
	"encoding/json"
	"github.com/gorilla/mux"
	currency "github.com/saurabhsisodia/gRPC/protos/currency"
)

type logger struct{
	l *log.Logger
	cc currency.CurrencyClient
}

func NewLogger(l *log.Logger , c currency.CurrencyClient)logger{
	return logger{l,c}
}

func (logger logger)ServeHTTP(rw http.ResponseWriter,r *http.Request){
	if r.Method==http.MethodGet{
		logger.l.Println("Serving HTTP-GET ...")
		enc:=json.NewEncoder(rw)
		mp:=mux.Vars(r)
		id,ok:=mp["id"]
		if ok{

			// calling gRPC server function running on localhost:9092
			rr:=&currency.RateRequest{
				Base:currency.CurrencyValue(currency.CurrencyValue_value["a"]),
				Destination:currency.CurrencyValue(currency.CurrencyValue_value["e"]),
			}
			prod:=data.GetProduct(id)
			resp,err := logger.cc.GetRate(context.Background(),rr)
			if err!=nil{
				logger.l.Println(err)
				return
			}
			prod.Price = prod.Price * resp.Rate


			err =enc.Encode(prod)
			if err!=nil{
				fmt.Fprintln(rw,"Error in encoding products")
			}
			return
		}
		err:=enc.Encode(data.GetProducts())
		if err!=nil{
			fmt.Fprintln(rw,"Error in encoding products")
			return
		}
	}
	if r.Method==http.MethodPost{
		fmt.Println(r.Body)
		prod:=&data.Product{}
		logger.l.Println("Serving HTTP-POST ...")
		err:=json.NewDecoder(r.Body).Decode(prod)  // always pass address in Decode
		if err!=nil{
			http.Error(rw,"error in decoding",http.StatusBadRequest)
			return

		}
		defer r.Body.Close()
		data.AddProduct(prod)
		logger.l.Printf("product %#v added succesfully\n",*prod)
		fmt.Fprintf(rw,"%#v",*prod)

	}
	if r.Method==http.MethodPut{
		logger.l.Println("Serving HTTP-PUT ...")
		mp:=mux.Vars(r)
		id,ok:=mp["id"]
		if !ok{
			http.Error(rw,"Path parameter required",http.StatusBadRequest)
			return
		}
		prod:=&data.Product{}
		err:=json.NewDecoder(r.Body).Decode(prod)
		prod.Id=id
		err=data.UpdateProduct(*prod)
		if err!=nil{
			http.Error(rw,"product does not exist",http.StatusNotFound)
			return
		}
		fmt.Fprintf(rw,"product %#v updatd successfully",*prod)

	}

	if r.Method==http.MethodDelete{
		logger.l.Println("Serving HTTP-DELETE ...")
		mp:=mux.Vars(r)
		id,ok:=mp["id"]
		if !ok{
			http.Error(rw,"Path Parameter required",http.StatusBadRequest)
		}
		err:=data.DeleteProduct(id)
		if err!=nil{
			http.Error(rw,"product does not exist",http.StatusNotFound)
			return
		}
		fmt.Fprintln(rw,"product deleted successfully")
	}
}