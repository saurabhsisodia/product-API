package main

import (
	"fmt"
	"time"
	"net/http"
	"log"
	"os"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
)
var mySigningKey = []byte("ILOVEYOU")
type myLogger struct{
	logger *log.Logger
}

// define a constructor for myLogger
func NewMyLogger(l *log.Logger) *myLogger{
	return &myLogger{logger:l}
}

func main(){

	l := log.New(os.Stdout,"[INFO] ",log.LstdFlags)
	Logger:=NewMyLogger(l)

	http.HandleFunc("/jwt",Logger.getJWT)
	log.Fatal(http.ListenAndServe(":8080",nil))  // nil for default servemux
}

func (logger *myLogger) getJWT(rw http.ResponseWriter,r *http.Request){
	logger.logger.Println("Serving HTTP-GET for JWT-TOKEN")
	jwtToken,err:=generateJWTToken()
	handleError(err)


	client := &http.Client{}
	req,_ := http.NewRequest("GET","http://localhost:9092/home",nil)
	req.Header.Set("Token",jwtToken)
	res,err := client.Do(req)
	handleError(err)
	body,err := ioutil.ReadAll(res.Body)
	handleError(err)
	defer res.Body.Close()

	fmt.Fprintln(rw,string(body))
}
func handleError(err error){
	if err!=nil{
		log.Fatal(err)
	}
}
func generateJWTToken()(string,error){
	token := jwt.New(jwt.SigningMethodHS256)   // create a new token struct instance

	claims := token.Claims.(jwt.MapClaims)  // type assertion, if dynamic type passes,
											//the dynamic value returns

	claims["authorized"] = true
	claims["user"] = "Saurabh Sisodia"
	claims["exp"] = time.Now().Add(time.Minute*10).Unix()

	tokenString,err := token.SignedString(mySigningKey)   // Complete token signed with Sinature
	if err!=nil{
		fmt.Errorf("%s",err.Error())
		return "",err
	}

	return tokenString,nil
}