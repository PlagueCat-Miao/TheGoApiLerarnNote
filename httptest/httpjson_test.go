package httptest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func init() {
	fmt.Println("http 各项测试开始 GET POST POST-json")

}


func TestHTTPGet(t *testing.T) {
	t.Log("Hellcat: http Get 请求测试")
	url := "http://127.0.0.1:8080/user/kkkk/dfg"

	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {

		panic(err)

	}

	defer resp.Body.Close()

	fmt.Println("GET 回复如下：" )
	fmt.Println("response Status:", resp.Status)

	fmt.Println("response Headers:", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response Body:", string(body))

	t.Log("Hellcat:  Get 请求测试 完成")
}

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required" `
	Password string `form:"password" json:"password" xml:"password" binding:"required" `
	Targetdb string `form:"targetdb" json:"targetdb" xml:"targetdb" binding:"required"`
	Dhash string `form:"Dhash" json:"Dhash" xml:"Dhash" binding:"required"`
	Status string `form:"status" json:"status" xml:"status" binding:"required"`
	Ip string `form:"ip" json:"ip" xml:"ip" binding:"required"`
	Capacity string `form:"capacity" json:"capacity" xml:"capacity" binding:"required"`
	Remain string `form:"remain" json:"remain" xml:"remain" binding:"required"`

}
func TestPostjson(t *testing.T) {
	t.Log("Hellcat: http Post-json 请求测试")
	request := Login{"hellcat","12345678","cloud","QmABC","3","127.0.0.1","400","300"}
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(request)


	url := "http://127.0.0.1:8080/do/test1"

	req, err := http.NewRequest("POST", url, requestBody)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {

		panic(err)

	}

	defer resp.Body.Close()

	fmt.Println("POST-json 回复如下：" )
	fmt.Println("response Status:", resp.Status)

	fmt.Println("response Headers:", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response Body:", string(body))


	ans :=new(Answer)
	//var ansP *Answer

	ans.AnswerResponse(string(body));
	fmt.Println("我是",ans.Dhash , "于" ,ans.ClientIP )

	t.Log("Hellcat:  Post-json 请求测试 完成")
}
type Answer struct {

	ClientIP string `json:"ClientIP"`
	code int `json:"code"`
	success bool  `json:"success"`
	Dhash string  `json:"Dhash"`

}

func (ans *Answer)AnswerResponse(input string) {


	if err := json.Unmarshal([]byte(string(input)), &ans); err == nil {
		fmt.Println("反序列化完毕: ",ans.ClientIP,"-",ans.Dhash)

	} else {
		fmt.Println("反序列化失败")
		fmt.Println(err)

	}

}
