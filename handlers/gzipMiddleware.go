package handlers
import (
	"net/http"
	"compress/gzip"
	"strings"
)

type GzipHandler struct{}


func (g *GzipHandler)GzipMiddleware(next http.Handler)http.Handler{
	return http.HandlerFunc(
		func (rw http.ResponseWriter,r *http.Request){

			// Header is a map[string][]string
			if strings.Contains(r.Header.Get("Accept-Encoding"),"gzip"){

				wrw:=NewWrappedResponseWriter(rw)
				wrw.rw.Header().Set("Content-Encoding","gzip")
				next.ServeHTTP(wrw,r)
				defer wrw.Flush()
				return

			}

			next.ServeHTTP(rw,r)
		})
}

type WrappedResponseWriter struct{
	rw http.ResponseWriter
	gw *gzip.Writer
}

func NewWrappedResponseWriter(rw http.ResponseWriter)*WrappedResponseWriter{
	gw:=gzip.NewWriter(rw)
	return &WrappedResponseWriter{rw:rw,gw:gw}
}

func (wr *WrappedResponseWriter) Header()http.Header{
	return wr.rw.Header()
}
func (wr *WrappedResponseWriter) Write(d []byte)(int,error){
	return wr.gw.Write(d)
}

func (wr *WrappedResponseWriter) WriteHeader(statuscode int){
	wr.rw.WriteHeader(statuscode)
}

func (wr *WrappedResponseWriter) Flush(){
	wr.gw.Flush()
	wr.gw.Close()
}