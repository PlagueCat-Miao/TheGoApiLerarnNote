package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

//不大写，不public
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

func httpGet(requestUrl string) (err error) {
	Url, err := url.Parse(requestUrl)
	if err != nil {
		fmt.Printf("requestUrl parse failed, err:[%s]", err.Error())
		return
	}

	params := url.Values{}
	params.Set("param","alice")
	params.Set("hello","cat")
	Url.RawQuery = params.Encode()

	requestUrl = Url.String()
	fmt.Printf("requestUrl:[%s]\n", requestUrl)
	
	client := &http.Client{}
	requestGet, _:= http.NewRequest("GET", requestUrl, nil)

	//requestGet.Header.Add("query", "googlesearch")
	//requestGet.Header.Add("content", "golang")

	resp, err := client.Do(requestGet)
	if err != nil {
		fmt.Printf("get request failed, err:[%s]", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp status code:[%d]\n", resp.StatusCode)
	fmt.Printf("resp body data:[%s]\n", string(bodyContent))
	return
}


func httpPost(requestUrl string) (err error) {
	data :=GetLogin();

	jsonData, _ := json.Marshal(data)
	fmt.Printf("requestUrl:[%s]\n", requestUrl)

	resp, err := http.Post(requestUrl, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		fmt.Printf("get request failed, err:[%s]", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp status code:[%d]\n", resp.StatusCode)
	fmt.Printf("resp body data:[%s]\n", string(bodyContent))
	return
}
func GetLogin() *Login{
	return &Login{
		User:     "hellcat",
		Password: "12345678",
		Targetdb: "cloud",
		Dhash:    "QmABCD",
		Status:   "3",
		Ip:       "127.0.0.1",
		Capacity: "400",
		Remain:   "300",
	}
}


func TestHttpCline(t *testing.T) {
	t.Log("API客户端: get post-json测试")

	Convey("[TestHttpCline] API客户端测试", t, func() {
		Convey("[Http Get] success", func() {
			var url = "http://127.0.0.1:8080/user/miaomiao"
			err := httpGet(url)
			So(err, ShouldBeNil)
			})
		Convey("[Http Post-json] success", func() {
			var url = "http://127.0.0.1:8080/do/test1"
			err := httpPost(url)
			So(err, ShouldBeNil)
		})

	})
	t.Log("API客户端: 测试完成")
}

