# GIN Framework 기본 설정
- restful 개발에 앞서 golang의 인기 있는 Web Framework인 gin을 살펴보도록 하자.

## go module 생성 및 gin 설치

### Go module 생성
- potato/simple-rest 이름의 go module을 생성한다.
```shell
❯ go mod init potato/simple-rest
```

### gin 설치
- Gin 설치 (https://github.com/gin-gonic/gin#installation)
```shell
go get -u github.com/gin-gonic/gin
```

## 간단한 Router 작성
- app이라는 이름의 디렉토리를 생성하고 main.go 파일을 생성한다.
- go는 자체 http 서버가 있기 때문에 gin을 router라고도 부른다.
- gin은 default 라우터를 제공하며, .GET/.POST와 같은 http method와 해당 경로로 요청을 받을 때 처리하는 Handler 를 연결한다.
```go
func newServer() *gin.Engine {
    r := gin.Default()
    r.GET("", helloHandler)
    r.GET("/:name", getOneSimpleHandler)
    r.POST("/", createSimpleHandler)

    return r
}
```

- handler는 gin.HandlerFunc와 동일한 형태면 가능하다.
```go
// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)
```
- Handler 작성한다. 실제 동작 내용은 다음시간에 하고 이번에는 Bootstrap 과정을 살펴보려한다.
```go
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
}

func getOneSimpleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
}

func createSimpleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
}
```
- Main함수를 작성하고, gin라우터를 실행한다.
```go
func main() {
	if err := newServer().Run(); err != nil {
		log.Fatalf("Could not run HTTP Server with (%v)", err)
		return
	}
}
```

## 실행
- go run을 이용해 실행한다.
```shell
❯ go run app/main.go
```

- 간단한 실행을 통해 우리가 설정하지 않은 포트(default:8080)으로 실행됨을 확인할 수 있다.
```shell
❯ go run app/main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.helloHandler (3 handlers)
[GIN-debug] GET    /:name                    --> main.getOneSimpleHandler (3 handlers)
[GIN-debug] POST   /                         --> main.createSimpleHandler (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```

- 다수의 Warning이 발생하는데 하나씩 살펴 보도록 하자.

- 제일 처음 나오는 Waring은 gin을 default로 사용할 경우 logger와 recovery가 default로 포함되어있다는 안내(?) 메시지 이다. - 굳이 warning으로 표현해야할까 싶다.
```shell
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
```

- 보기 흉하니 Default로 사용하지 gin.New()를 통해 코드를 변경해보자.
```go
func newServer() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("", helloHandler)
	r.GET("/:name", getOneSimpleHandler)
	r.POST("/", createSimpleHandler)

	return r
}
```

- 두번째 Warning은 Running Mode에 대한 것으로 golang이 build형 언어이다 보니, debug/release Mode에 따라 실행환경을 구분할 수 있다.
- 현재 GIN이 debug 모드로 실행 중이라는 안내(?) 메시지 이다. release mode를 사용하면 gin 자체에서 출력되는 로그 출력을 제외 할 수 있다.
- 친절하게도 환경변수와 Code를 통해 해당 내용을 반영할 수 있다. (추후 배포단계에서 dockerized 할때 변경하도록 하고 지금은 debug 를 사용하자. )
```go
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)
```

- 마지막 Warning은 X-Forwarded-For Header와 관련이 있다. 
- 운영환경에서 Backend Server로 인입되는 호가 중간에 API G/W나 WebServer등을 거칠때 EndUser의 IP를 전달 하기 위해 X-Forwarded-For 헤더를 사용해 전달한다.
- Gin은 default로 모든 요청을 신뢰하지만, 운영환경에서 backend server는 모든 요청이 EndUser가 아닐 수 있어 신뢰 할 수 있는 G/W등 에서만 X-Forwarded-For를 사용하도록 한다.
```shell
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
````

- SetTrustedProxies 설정을 통해 X-Forward-For Header를 사용할 수 있다.
- 설정은 IP를 등록하거나
```go
if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
    return nil
}
```
- CIDR을 등록 할 수 있다.
```go
if err := r.SetTru****stedProxies([]string{"127.0.0.0/8"}); err != nil {
    return nil
}
```

- 만일 CDN을 사용하는 경우 X-Forward-For 가 아닌 별도의 header를 사용할 수 있기 때문에 별도로 Header를 설정 할 수 있다.
```go
r.TrustedPlatform = gin.PlatformCloudflare
r.TrustedPlatform = gin.PlatformGoogleAppEngine
r.TrustedPlatform = "X-POTATO-IP"
```

- 이제 모든 Debug Mode를 제외한 모든 Warning이 사라졌고, 다음과 같은 실행 화면을 볼수 있다.
```shell
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.helloHandler (3 handlers)
[GIN-debug] GET    /:name                    --> main.getOneSimpleHandler (3 handlers)
[GIN-debug] POST   /                         --> main.createSimpleHandler (3 handlers)
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```

- 하지만 마지막으로 하나만 더 살펴 보자. 
- Warning은 아니지만 "[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default" 가 찝찝해 보인다.  
(이런걸 Warning으로 띄워줘야 하지 않나 싶다.)
- default로 8080포트를 사용하지만 별도의 Port를 지정해 줄 수 있다.
```go
if err := newServer().Run(":8080"); err != nil {
    log.Fatalf("Could not run HTTP Server with (%v)", err)
    return
}
```