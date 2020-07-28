FROM golang:latest

WORKDIR /go

RUN go get github.com/gin-gonic/gin
RUN go get github.com/codegangsta/gin
RUN go get github.com/go-sql-driver/mysql

WORKDIR /
RUN mkdir go_work

COPY ./go_work/default /go_work/default
WORKDIR /go_work/default

CMD [ "gin" , "-i" , "--all" , "run" , "main.go"]
