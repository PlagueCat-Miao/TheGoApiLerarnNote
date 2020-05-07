package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/PlagueCat-Miao/TheGoApiLerarnNote/GINtest"

	"net/http"

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




	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(ctx context.Context){


	}(ctx)

	var n  Name;
	n1 := new(Name)
	var n2,n3 *Name
	n2=&n;
	n3=n1;
    n2.name="hell"
    n3.name="cat"

	fmt.Println("n",n,"n1:",n1)
	var fn,fn2  Name;
	input(fn)
	inputP(&fn2)
	fmt.Println("n",fn.name,"n1:",fn2.name)

	//router.Run(":8080")


}
func input(a Name){
	a.name= "k"
}
func inputP(a* Name){
	a.name= "kP"
}