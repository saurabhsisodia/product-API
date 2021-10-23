package main

import (
	"log"
	"net/http"
	"REST-API/handlers"
	"REST-API/data"
	"os"
	
	"github.com/gorilla/mux"
	gorillaHandler "github.com/gorilla/handlers"  // for Implementing CORS
	currency "github.com/saurabhsisodia/gRPC/protos/currency"
	grpc "google.golang.org/grpc"

)

func main(){	
	gzipHandler:=&handlers.GzipHandler{}
	base:=`c:\\users\astrea\Desktop\REST-API`
	
	// create gRPC client
	conn,err:=grpc.Dial("localhost:9092",grpc.WithInsecure())  // in production used TLS
	if err!=nil{
		panic(err)
	}
	defer conn.Close()
	cc:=currency.NewCurrencyClient(conn)


	l:=log.New(os.Stdout,"Product-API ",log.LstdFlags) 
	logger:=handlers.NewLogger(l,cc)

	router:=mux.NewRouter()  // top level router
	// GET request Subrouter
	getRouter:=router.Methods(http.MethodGet).Subrouter()
	getRouter.Handle("/products",logger)
	getRouter.Handle("/product/{id:[0-9]+}",logger)
	getRouter.Handle("/{[a-z]+\\.[a-z]{3}}",http.FileServer(http.Dir(base)))   // serving static files
	getRouter.Use(gzipHandler.GzipMiddleware)  // middleware for Gzip Compression if Accept-Encoding is "gzip"

	//POST request Subrouter
	postRouter:=router.Methods(http.MethodPost).Subrouter()
	postRouter.Handle("/products",logger)
	postRouter.Use(data.ValidationMiddleware)   // adding Middleware for json field validation

	//PUT request Subrouter
	putRouter:=router.Methods(http.MethodPut).Subrouter()
	putRouter.Handle("/product/{id:[0-9]+}",logger)
	putRouter.Use(data.ValidationMiddleware)    // adding Middleware for json field validation

	// DELETE request Subrouter
	delRouter:=router.Methods(http.MethodDelete).Subrouter()
	delRouter.Handle("/product/{id:[0-9]+}",logger)




	/*
		CORS = Cross Origin Resource Sharing
		if origins of two apps are different then to share resources the server has to implement CORS
		mechanism and must set its headers like
			Access-Control-Allowed-Origin
			Access-Control-Allowed-Headers
			Access-Control-Allowed-Methods
					etc ..
	*/
	cors:=gorillaHandler.CORS(gorillaHandler.AllowedOrigins([]string{"*"})) // set some headers

	//can also allow all domains or methods or ports or subdomain
	// for all domains, by setting []string{"*"}

	log.Fatal(http.ListenAndServe(":8080",cors(router)))  // wrap the top level router with cors
}
