FROM golang:1.19.4-alpine3.17

WORKDIR /app

ENV GIN_MODE=release

#Download dependency
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .


RUN go build  -o /main ./cmd/app/main.go

EXPOSE 8000

CMD [ "/main" ]