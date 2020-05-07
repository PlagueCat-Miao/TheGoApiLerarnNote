package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/plaguecat-miao/TheGoApiLerarnNote/GINtes"
)

type  Name struct {
	name string
}

func init() {
	fmt.Println("gin与协程，，请手动退出")
}

func main() {

	router := gin.Default()
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name) //以字符串的形式回复
	})
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	router.POST("/do/test1", GINtest.Handle1)




	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	/*go func(ctx *context.Context){


	}(ctx)
*/
	var n Name;
	n1 := new(Name)
	var n2,n3 *Name
	n2=&n;
	n3=n1;
    n2.name="hell"
    n3.name="cat"

	fmt.Println(n,n1)

	//router.Run(":8080")


}
