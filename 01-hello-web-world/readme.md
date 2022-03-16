# Hello Web World

- 간단한 라우터를 만들어서 Go의 내장 웹서버에 연결시킨다.

## 1. Go 프로젝트 생성

- 적절한 directory를 만들어서 명령어를 통해 Go 프로젝트를 생성한다.

```sh
❯ mkdir 01-hello-web-world
❯ cd 01-hello-web-world
```

- go cli를 이용해서 초기화한다. init 뒤의 모듈명은 해당 모듈을 라이브러리로 사용할 때 이용되기도 한다.

```sh
go mod init potato/hello-web-world
```

- 라우터 기능을 제공하는 gin을 설치한다. 
- go get은 maven이나 gradle 처럼 외부의 라이브러리를 가져오는 기능을 제공한다.
- gin은 spring과 같은 framework으로 생각할 수 있다.

```sh
go get -u github.com/gin-gonic/gin
```

## 2. Main 코드 생성

- /application/main.go 파일을 생성하고 아래의 내용은 입력한다.
  
```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Web World",
		})
	})
	r.Run()
}

```

- 하나씩 살펴보면 Go의 Entrypoint 인 main() 을 포함하는 Package 이다.

``` go
package main
```

- go get으로 다운로드 받은 gin 라우터를 import 한다.

```go
import "github.com/gin-gonic/gin"
```

- java와 달리 go는 언어차원에 내장 웹서버(net/http)를 포함하고 있으며, 우리가 go get을 다운로드 받은 gin 라우터가 내장 웹서버를 사용한다.

```go
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Web World",
		})
	})
	r.Run()
}
```

- main() 함수는 프로그램을 시작하는 지점이고, gin이 r 이라는 router 객체를 생성하고 GET Method를 이용하는 "/" Path의 실제 동작인 function을 정의 했다.
- 물론 function은 외부로 추출할 수 있다.

## 3. 실행
```sh
❯ go run application/main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```

