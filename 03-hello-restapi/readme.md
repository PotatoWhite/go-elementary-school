# Simple Rest API 개발
- GET, POST 의 Request 방법과 Params의 획득방법을 알아본다.

## DTO 등록

### Struct + annotation


### QueryString -> 
- potato/simple-rest 이름의 go module을 생성한다.

### RequesBody -> ShouldBind
- potato/simple-rest 이름의 go module을 생성한다.

### Content-Type

### gin.H를 이용하 HashMap Response

### annotation을 이용해 Struct -> JSON 변환시 이름을 변경해 보자
```go
type Simple struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
```