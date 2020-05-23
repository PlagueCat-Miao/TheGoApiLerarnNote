package GINtest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"log"
	"net/http"
	"testing"
)

//不大写，不public
type Login2 struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required" `
	Password string `form:"password" json:"password" xml:"password" binding:"required" `
	Targetdb string `form:"targetdb" json:"targetdb" xml:"targetdb" binding:"required"`
	Dhash string `form:"Dhash" json:"Dhash" xml:"Dhash" binding:"required"`
	Status string `form:"status" json:"status" xml:"status" binding:"required"`
	Ip string `form:"ip" json:"ip" xml:"ip" binding:"required"`
	Capacity string `form:"capacity" json:"capacity" xml:"capacity" binding:"required"`
	Remain string `form:"remain" json:"remain" xml:"remain" binding:"required"`

}

func init() {
	fmt.Println("Get - key value")
}

func TestGet(t *testing.T) {
	t.Log("Hellcat: GIN 请求测试")

	router := gin.Default()

	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	//可以获取/:xxxx处的路径名称，并跳转至/xxxx
	router.GET("/user/:name", GetGetReq)
	router.POST("/do/test1", PostHandle)


	t.Log("Hellcat: 请手动退出")
	fmt.Print("Hellcat: 请手动退")
	router.Run(":8080")


}



func GetGetReq(c *gin.Context) {
	name := c.Param("name")
	hello := c.Query("hello")//params
	c.Set("RequestContext","由服务端设置，不存在于http中") // context
	_ ,exite := c.Get("RequestContext")
	fmt.Printf("  -  Context中的全局kv: RequestContext--%v \n      一般Get与Set是用于检查是否加载了中间件的\n",exite)
	c.String(http.StatusOK, "路径：%s 参数hello： %v", name,hello) //以字符串的形式回复
}

func PostHandle(c *gin.Context) {

	var json4 Login2
	if err := c.ShouldBindJSON(&json4); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ClientIP := c.ClientIP()

	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"targetdb": json4.Targetdb,
		"Dhash":    json4.Dhash,
		"status":   json4.Status,
		"ip":       json4.Ip,
		"ClientIP": ClientIP,
		"capacity": json4.Capacity,
		"remain":   json4.Remain,
		"success":  true,
	})

}
