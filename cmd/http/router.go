package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRouter(r *gin.Engine) {
	r.GET("/health_check", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	r.POST("/file/upload", FileUpload)
	r.GET("/sts", StsToken)
	r.GET("/file/download/:media_id", FileDownload)
	r.GET("/file/info/:media_id", FileDownload)
}
