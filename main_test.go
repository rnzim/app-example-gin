package main

import (
	"gin-api/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func RoutesConfig()  *gin.Engine {
	r:= gin.Default()
	return r
}


func TestRoutes(t *testing.T) {
   r := RoutesConfig()

   r.GET("/",controllers.Index)
   req,_ := http.NewRequest("GET","/",nil)
   resp := httptest.NewRecorder()
   r.ServeHTTP(resp,req)

   if resp.Code != http.StatusOK{
	t.Fatalf("Status Error")
   }

}