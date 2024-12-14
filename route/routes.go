package route

import "github.com/gin-gonic/gin"

func HandleRequest(){
	r := gin.Default()
	defer r.Run()
}