# TheGoApiLerarnNote
主要学习 gin框架与 net/http
## 下载
``` 
    cd $GOPATH/src/github.com/PlagueCat-Miao/TheGoApiLerarnNote;git pull origin master;git diff  

```
## 上载
```
    cd $GOPATH/src/github.com/PlagueCat-Miao/TheGoApiLerarnNote;git add . ;git commit -m "快速上传"; git push origin master 

```
## 墓碑
### gin
 - [Gin框架中文文档](https://www.jianshu.com/p/98965b3ff638/)
 - [Go语言Web框架--Gin介绍和使用](https://blog.csdn.net/qq_34777600/article/details/81160167)
#### json (BindXXX)
  当我们使用 content-type:application/json 报文格式发送json格式数据时，不能用PostForm
  而是要用BindJSON/ShouldBindJSON,如下:
  ``` go-gin-Handler-json 
        type Login struct {
            User     string `form:"user" json:"user" xml:"user"  binding:"required"`
            Password string `form:"password" json:"password" xml:"password" binding:"required"`
        }
        ...
        func xxxx  { ...
            var json Login
            if err := c.ShouldBindJSON(&json);
        } 
  ```
- [Gin框架中文文档](https://www.jianshu.com/p/98965b3ff638/)
- [gin的BindJSON和ShouldBindJSON，ShouldBindWith的区别](https://blog.csdn.net/heart66_A/article/details/100796964)
#### context 核心介绍
 - [gin源码阅读之三 -- gin牛逼的context](https://www.jianshu.com/p/73bf8fe7a745)

### go
#### 获取go变量
 - go env 获取环境变量
 - go env XXXX 可以干净整洁地输出需要的变量值（如： `go enc GOOS `）
 - [go env](https://wiki.jikexueyuan.com/project/go-command-tutorial/0.14.html)
#### go build /install
 - go build 用于生成可执行文件
 - go install 用于(无main)生成库、工具、(有main)可执行文件; 主要是编译二进制文件）
 - [深入理解 go build 和 go install](https://www.jianshu.com/p/3db831d9b553)
#### go 私有/共有变量
众所周知 go的struct类型中，首字母小写的变量为私有（private）变量
 - 跨文件不能直接访问/赋值
 - 同文件其他函数不能直接访问/赋值
 - 同文件其他函数使用地址符&获取数组名后，仍不能访问/赋值！！如：
  `var json4 Login; err := c.ShouldBindJSON(&json4);` 时，json4中的小写变量将不会被覆盖，且程序不会报错；
  

### postman
#### 如何卸载
 - 删除 C:\Users\Lenovo\AppData\Roaming和C:\Users\lenovo\AppData\Local 目录下。
 - [参考win10删除Postman](https://www.jianshu.com/p/27842e040678)
#### 发送json
- Collection - request - Hearder - +content-type~application/json
- Collection - request - Body - raw - +{xxjsonxx}
- send
- [postman 发送json请求](https://blog.csdn.net/weixin_37569048/article/details/81456561)

#### http
#### cookie 的两个流程：request.cookie \Response.Cookies
 - response.cookie：用于在客户端写入cookie值。若指定的cookie不存在，则创建它。若存在，则将自动进行更新。结果返回给客户端浏览器。
 - request.cookie：设置cookie的最大有效期为30天，然后通过Response回送cookie到浏览器。
 - [response.cookie和request.cookie的区别 ](https://www.cnblogs.com/try-chi/p/11995231.html)
#### 路由路径与请求方法
 - 同一路由路径，不同请求方法（如：GET\POST） 是两个请求
   - 不会冲突
   - 需要分别注册HandlerFunc
   - 发送的时候，自然是走各自的注册的路由，在需要的时候，可以绑定相同的处理函数（HandlerFunc）以达到类似“PHP-$_REQUEST超全局数组”的效果，如：
     ```
      router.GET("login", LoginCheck)
      router.POST("login", LoginCheck)
     ```
    -[gin框架相关参考](https://www.cnblogs.com/-beyond/p/9391892.html)
#### GET/POST
- GO
   ``` go GET 流程
        req, err := http.NewRequest("GET", url, nil)
        client := &http.Client{}
        resp, err := client.Do(req)
        defer resp.Body.Close()
   ``` 
   ``` go POST-json 流程
        ...
        //requestBody内容为json
        req, err := http.NewRequest("POST", url, requestBody)
        resp, err := client.Do(req)
  
        //反序列化
        var ans Answer
        json.Unmarshal([]byte(string(resp)), &ans)
   ```
- curl
  -  POST 为例子
    ``` liunx                                                                               
            curl -X POST http://localhost:8080/loginJSON -H 'content-type:application/json' -d '{"user":"manu","password":"123"}'
    ```

### makefile
#### makefile中使用 shell
 - 正常指令直接输入如：`go build`
 - 需要Makefile读取shell命令输出时： `VAL=$(shell xxxx)`
 - 前后有关联的语句需要在同一行以进入同一进程 ： 
 
              ``` 
                cd xxxx \(回车)
                go build 
              ```
  - 使用Makefile变量 `$(VAL)` 使用shell系统变量 `$$VAL`
  - [总结-makefile中的shell调用---注意事项](https://blog.csdn.net/frank_jb/article/details/81708832)
  - [示例-获取文件所在路径](https://www.cnblogs.com/catgatp/p/6527243.html)



## 雷区
### 1.修改对象时，需要使用指针
 - `func (a *Ans)xxxx {` 
     - 传递了指针后，函数内内容可不比追究是否使用指针格式（不区分`a = a+b`or `*a =*a + *b`)
     - 不使用指针，如：`func (a Ans)xxxx {` 传递的是副本，即使使用地址`&a` 也不会影响函数外的struct（即调用该函数的对象） 
### 2.新建类
  - `var ans Answer`  对象
  - `ansP := new(Answer)`  指针，有对象
  - `var ansP *Answer` 指针 
    - `ansP.name ` 指针调用成员时不用解指针（`*a`） 没指时用 就Panic
    - `ansP = &ans`当指针用  没指时用 就Panic
  - `requestBody := new(bytes.Buffer)` bytes.Buffer里面是切片，切片都是指针
  - `func input(a Answer)`  a是副本，出函数 a不会变化
  - `func input(a *Answer)`  &a是地址指针，出函数后 a会变化
  - `func(ctx context.Context)` 首先ctx 是副本，但是通道是指针：名相同，就是公用的，变化所有人都有反应（指竞争关系）
### 3.xxxx_test.go 
  - 整个文件是不能和此包一起被import的
  - go test -v xxxx_test.go 查看详情
  - 对于循环有限时（我运行了600s，test自动FAIL退出）
### 4. go get /mod 大小写 （代理：On）
  - 注意大小写 go get xxx 遇到大写时会替换成！小写 （如：`PLagueCat-Miao => !plague!cat-!miao`）但是import 和 build 路径时又不支持`!`
    - 不要改 pkg文件夹下的文件如：go.mod 它和go.sum 他们有关联 单改没有用
    - import 当然支持大写 如：`github.com/PlagueCat-Miao/TheGoApiLerarnNote`
       - 虽然被保存成 `github.com/!plague!cat-!miao/!the!go!api!lerarn!note`
       - 但是内部(见go.mod）仍正常大写，可见go.mod才是正确的检索依据
       - 当你在因个新项目创建go.mod `go mod init`时，注意看module
          - 她根是据此时的$GOPATN路径自动分析处本包名字
          - 她区分`github.com/!plague!cat-!miao/!the!go!api!lerarn!note`
          - 她区分`github.com/plaguecat-miao/thegoapilerarnnote`
          - 她区分`github.com/PlagueCat-Miao/TheGoApiLerarnNote`
          - 以上三个都是严格区分的
       - 所以你要**使用大写**的话，
          - 新项目：请在将src文件系统路径统一大写后(并去掉!)，再次`go init`
          - pkg包：不用管pkg下文件系统路径，import该大写就大写
    - 当你go get      
        - `github.com/PlagueCat-Miao/TheGoApiLerarnNote`
           - 可以get，得到文件系统 `github.com/!plague!cat-!miao/!the!go!api!lerarn!note`
        - `github.com/plaguecat-miao/TheGoApiLerarnNote`
           - 可以get，得到文件系统 `github.com/plaguecat-miao/!the!go!api!lerarn!note`
        - `github.com/plaguecat-miao/thegoapilerarnnote`         
           - 可以get，得到文件系统 `github.com/plaguecat-miao/thegoapilerarnnote`
        - import
            1. 大写对应文件系统路径为 pkg/mod/!xxxx(替换成！+小写)
            2. go.mod 
              - 如果项目有，内容是原版，且 import 与 go.mod 的module 必须一致(要是的module大写你还小写import 就炸)
              - 如果没有go.mod  满足上文的路径就可以（不过pkg 都是go get 自动下和命名文件路径 ，所以小写大写无所谓）
       - 你不想走 pkg 相当于走github库(pkg/mod)
           -  GO111MODULE=off 切换为本地模式 （然后你就被祖国的墙把go get墙住
       - 所以 ！只应该出现在 pkg 下，作为自动命名，src下请正常命名
    
    
### 5. go env -w GO111MODULE=off 代理
  - 没办法代理会阻止你使用本地路径寻找包（哪怕就在自己目录下）
    - 报错：` cannot find module for path xxxxxxxx`
    - 请使用github 上传后 在下载的方式