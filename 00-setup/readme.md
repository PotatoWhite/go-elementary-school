# Go 프로젝트 준비사항

- Golang 설치
- Git 설치

## 1. Golang 설치 관련

- 공식 설치 지침 (<https://go.dev/doc/install>) : 잘 보고 설치 한다.
  
### 1.1 다운로드 GO

- 각자 개발환경에 맞게 1.17.8 버전 다운로드(<https://go.dev/dl>)

    ```sh
    ❯ wget https://go.dev/dl/go1.17.8.linux-amd64.tar.gz
    --2022-03-15 23:05:18--  https://go.dev/dl/go1.17.8.linux-amd64.tar.gz
    Resolving go.dev (go.dev)... 216.239.32.21, 216.239.34.21, 216.239.36.21, ...
    Connecting to go.dev (go.dev)|216.239.32.21|:443... connected.
    HTTP request sent, awaiting response... 302 Found
    Location: https://dl.google.com/go/go1.17.8.linux-amd64.tar.gz [following]
    --2022-03-15 23:05:19--  https://dl.google.com/go/go1.17.8.linux-amd64.tar.gz
    Resolving dl.google.com (dl.google.com)... 142.251.42.142, 2404:6800:4004:821::200e
    Connecting to dl.google.com (dl.google.com)|142.251.42.142|:443... connected.
    HTTP request sent, awaiting response... 200 OK
    Length: 134902354 (129M) [application/x-gzip]
    Saving to: ‘go1.17.8.linux-amd64.tar.gz’

    go1.17.8.linux-amd64.tar.gz            100%[=========================================================================>] 128.65M  40.1MB/s    in 3.3s    

    2022-03-15 23:05:22 (38.4 MB/s) - ‘go1.17.8.linux-amd64.tar.gz’ saved [134902354/134902354]
    ```

- 압축 해제 및 설치
  
  ```sh
  ❯ sudo tar -C /usr/local -xzf go1.17.8.linux-amd64.tar.gz
  ```

- 환경 변수 등록 (.profile 등)
  - GOROOT : GO가 설치 된 경로 (JDK)
  - GOPATH : 외부 패키지가 설치 된 경로 (MAVEN)
  
  ```sh
  export GOROOT=/usr/local/go
  export GOPATH=$HOME/go
  export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
  ```

- 설치 확인

  ```sh
  ❯ go version
  go version go1.17.8 linux/amd64
  ```

## 2 install GIT

- 넘어가자.

## 3. 설치 확인 및 실행

- Go는 소스코드를 바로 실행 할 수도 있지만, 빌드하여 실행할 수도 있다.

- hello.go 파일을 생성해 아래의 내용을 작성한다.
  
  ```go
  package main
  
  import "fmt"
  func main(){
    fmt.Println("Hello World")
  }
  ```

- 실행(no build)
  
  ```sh
  ❯ go run ./hello.go
  Hello World
  ```

- 빌드 및 실행
  
  ```sh
  ❯ go build ./hello.go
  ❯ ./hello
  Hello World
  ```
  