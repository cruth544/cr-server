FROM golang:latest

RUN mkdir /app
ADD /src/. /app/
WORKDIR /app

# RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

RUN go build -o main .

ENTRYPOINT CompileDaemon --build="go build -o main ." --command=./app/main
# CMD ["/app/main"]

