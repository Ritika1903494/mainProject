package main


import (
	"github.com/gin-gonic/gin"
	control "mainproject/controller"
	cors "github.com/rs/cors/wrapper/gin"
 )

 func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/Main", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello gin",
		})
	})
	r.POST("/Main/CreateNotebook",control.CreateNote)
	r.GET("/Main/GetNotebooks",control.GetNotes)
	r.POST("/Main/CreateSection",control.CreateSection)
	r.GET("/Main/GetSections/:note_id",control.GetSections)
	r.POST("/Main/CreatePage",control.CreatePage)
	r.GET("/Main/GetPages/:section_id",control.GetPages)
	r.POST("/Main/CreateContent",control.CreateContent)
	r.GET("/Main/GetContent/:page_id",control.GetContent)
	r.POST("/Main/UpdateContent/:content_id",control.UpdateContent)
	r.Run() // listen and serve on 0.0.0.0:8080
}