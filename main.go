package main

import (
	"github.com/Protoculos/backgin/domain"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/:name", IndexHandler)
	router.Run()
}

func IndexHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	if c.Request.Header.Get("Content-Type") == "application/xml" {
		c.Writer.Header().Add("Content-Type", "application/xml")
		c.XML(200, domain.Person{
			FirstName: name,
			LastName:  "Ali",
		})
	} else {
		c.Writer.Header().Add("Content-Type", "application/json")
		c.JSON(200, domain.Person{
			FirstName: name,
			LastName:  "Ali",
		})
	}

}
