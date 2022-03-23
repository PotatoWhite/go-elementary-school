# 03-hello-restapi : Restful API를 작성할 떄 

- GET, POST 의 Request 방법과 Params의 획득방법을 알아본다.


## 1. Query String에 포함된 Parameter 획득방법

- Router의 정의 : QueryString은 별도의 정의가 없다. 어떤 method라도 받아서 처리할 수 있다.
```go
r.GET("/", simpleHandler)
```

- 파라미터의 획득 : Query를 사용한다.
```go
func simpleHandler(c *gin.Context) {
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%v 입니다.", name),
	})
}
```

- http 테스트 : test.http 를 참고하자.
```http request
### queryString
GET localhost:8080/?name=당근
```


## 2. URL Path에 포함된 Parameter 획득방법

- Router의 정의 : Path에 미리 "name" 파라미터를 받을 위치를 지정한다.
```go
r.GET("/:name", pathParamHandler)
```

- 파라미터의 획득 : Gin Router에 미리정해둔 이름 "name"을 Param gin.Context.Param 메소드를 이용해 획득한다. 
```go
func pathParamHandler(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%v 좋아요.", name),
	})
}
```

- http 테스트 : test.http 를 참고하자.
```http request
### path parameter
GET localhost:8080/당근
```

## 2-1. URL Path에 여러개가 들어오는 경우

- Router의 정의 : Path에 미리 "name", "quantity" 파라미터를 받을 위치를 지정한다.
```go
r.GET("/:name/:quantity", pathParamsHandler)
```

- 파라미터의 획득 : Gin Router에 미리정해둔 이름 "name"을 Param gin.Context.Param 메소드를 이용해 획득한다. 
```go
func pathParamsHandler(c *gin.Context) {
	stringValue := c.Param("name")
	numericString := c.Param("quantity")
	if quantity, err := strconv.Atoi(numericString); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%v, %v 개 주세요.", stringValue, quantity),
		})
	}
}
```

- http 테스트 : test.http 를 참고하자.
- 만일 ":name"이 아니고 숫자형인 ":height" 라면? -> 문자로 받아서 변환하다.
```http request
### path string/numeric
GET localhost:8080/당근/100
```

## 3. Request Body로 들어오는 경우

- 테스트를 위해 DTO로 사용할 struct 정의 해서 테스트한다.

### 3.1 Simple struct 정의

- 간단하게 Id, Name 두개의 Member만 정의 해보자, 이후 `json:"id"`로 보이는 내용은 json으로 변환 될때 이름으로 사용된다.(gin은 json->object 변환시는 대소문자 구분없이 잘 받아주기 때문에 object->json 변환시 이용된다.)
- "required"는 필수로 들어와야 하는 json 필드를 정의한다.

```go
package dto

type Simple struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}
```


### 3.2 테스트

- Post Method는 Request Body를 사용할 수 있다. 
```go
r.POST("/", requestBodyhandler)
```

- 파라미터의 획득 : ShouldBind를 통해 Json을 미리정의된 dto.Simple로 받을 수 있다.
- 이 시점에 required로 설정된 Json 필드가 없다면 에러가 발생한다.
```go
func requestBodyhandler(c *gin.Context) {
	var reqeustBody dto.Simple
	if err := c.ShouldBind(&reqeustBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("err: %v", err),
		})
		return
	} else {
		c.JSON(http.StatusOK, reqeustBody)
	}

	return
}
```

- http 테스트 : test.http 를 참고하자.
```http request
### create simple
POST localhost:8080
Content-Type: application/json

{
  "id": 123,
  "name": "당근"
}
```

