# 04-hello-test : HTTP Request를 통한 API 수준의 Test Code 작 

- GET, POST의 Request 방법과 간단히 테스트 코들 작성해 보자


## 0. Refactoring 

### 1. 운영환경에서 사용할 DTO를 추가적으로 작성한다.
```go
package dto

type BasicResponse struct {
	Code    int    `json:"code" binding:"required"`
	Message string `json:"message" binding:"required"`
}
```

## 2. Response Code를 통일 한다.
- Command 현, POST, PUT, DELETE 형은 BasicResponse 로 통일
- Query 형 , Single, Collection을 이용한 Object로 통일

## 3. Golang의 Unit Test Code

## 4. Coverage 확인