# 베이스 이미지로 공식 Golang 이미지를 사용합니다.
FROM golang:1.17.5-alpine3.15

# 작업 디렉토리를 설정합니다.
WORKDIR /app

# 애플리케이션의 의존성 파일을 복사합니다.
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# 애플리케이션 소스 코드를 복사합니다.
COPY . .

# 애플리케이션을 빌드합니다.
RUN go build -o /go/bin/app

# 컨테이너 시작 시 실행할 명령을 지정합니다.
CMD ["/go/bin/app"]