package main

import (
	"context"
	"fmt"
	"github.com/PlagueCat-Miao/TheGoApiLerarnNote/GINtest"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)


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

	router.GET("/do/get")

	server := &http.Server{
		Addr:              ":8080",
		Handler:           router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed{
			log.Fatalf("Listen:%s\n",err)
		}
	}()
	// 等待中断信号以优雅地关闭服务器（设置 10 秒的超时时间）
	fmt.Println("等待ctrl+c 启动Shutdown-Server")
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	//得到在当前上下文，并设定的了死亡时间
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	//首先停止接受所有新请求，并一个个处理旧请求，
	// 获取上下文，以获取死亡时间
	// 于min( 上下文的死亡时间，全部请求处理结束) 时间 停止堵塞。
	if err := server.Shutdown(ctx);err!= nil {
		log.Fatal("server shutdown: ",err)
	}

	log.Println("server exiting...")

	//router.Run(":8080")


}
