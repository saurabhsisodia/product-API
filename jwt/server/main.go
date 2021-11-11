package main
import(
	"fmt"
	"log"
	"os"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)
var mySigningKey = []byte("ILOVEYOU")
type myLogger struct{
	logger *log.Logger
}

func NewmyLogger(l *log.Logger)*myLogger{
	return &myLogger{l}
}

func main(){
	l:=log.New(os.Stdout,"[INFO] ",log.LstdFlags)
	Logger:=NewmyLogger(l)

	Logger.handleRequest()
}

func (logger *myLogger) handleRequest(){

	http.Handle("/home",logger.isAuthorized(logger.home))  // use middleware 
	log.Fatal(http.ListenAndServe(":9092",nil))
}

func (logger *myLogger) home(rw http.ResponseWriter,r *http.Request){
	logger.logger.Println("serving HTTP-GET request")
	fmt.Fprintln(rw,"very sensitive information")
}

func (logger *myLogger) isAuthorized(endpoint func(http.ResponseWriter,*http.Request)) http.Handler{
	return http.HandlerFunc(func(rw http.ResponseWriter,r *http.Request){

		if r.Header["Token"]!=nil{
			token,err:=jwt.Parse(r.Header["Token"][0],func (token *jwt.Token)(interface{},error){
				if _,ok:=token.Method.(*jwt.SigningMethodHMAC);!ok{
					return nil,fmt.Errorf("error during Method Handling in JWT")
				}
				return mySigningKey,nil
			})
			if err!=nil{
				logger.logger.Println(err.Error())
				fmt.Fprintln(rw,err.Error())
			}
			if token.Valid{
				logger.logger.Println("Token matched")
				endpoint(rw,r)
			}
		}else{
			logger.logger.Println("Token not Matched")
			rw.Write([]byte("Not Authorized"))
		}
	})
}