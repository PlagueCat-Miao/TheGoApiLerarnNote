package GINtest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func init() {
	fmt.Println("gin与http学习，test文件，请手动退出")
}



func TestGin(t *testing.T) {
	t.Log("Hellcat: GIN 请求测试")

	router := gin.Default()

	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	//可以获取/:xxxx处的路径名称，并跳转至/xxxx
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name) //以字符串的形式回复
	})

	// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
	// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
	//可以获取/:xxxx/*yyyy处的路径名称，并跳转至/xxxx
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.POST("/do/test1", Handle1)

	t.Log("Hellcat: 请手动退出")
	fmt.Print("Hellcat: 请手动退")
	router.Run(":8080")


}
