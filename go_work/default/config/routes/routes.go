package routes

import (
	"os"
    "net"
    "net/http/fcgi"
	"github.com/gin-gonic/gin"

	"default/controllers/top"
    "default/models/logs"
)

func Init() {

	// gin initialize
    router := gin.Default()
	router.LoadHTMLGlob("resources/html/*.html")
	
	// please write routing in then.
	router.GET("/", func(c *gin.Context) { top.Index( c ) })
	router.GET("/admin", func(c *gin.Context) { top.Admin( c ) })
	
	// fcgi access recieve
    l, err := net.Listen("tcp", os.Getenv("ACCESS_PORT") )
    logs.PutError( 1 , err )
    if err := fcgi.Serve(l, router); err != nil {
        logs.PutError( 1 , err )
    }
}