package router

import (
	c "../controller"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	//DefinePath
	orbit := r.Group("api/user/")
	{
		orbit.POST("DefinePath", c.DefinePath)
	}
}