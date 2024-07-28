#Specify the language version
#Work directory
#copy file for dependencies
#run to download dependencies
#then copy source code file to be run
#Check for Go modules and download dependencies if necessary
#which file to be run or run go build to run binary
#specify the port also for handling http

FROM golang:1.22.2 as builder

WORKDIR /app
COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o golang-wsl

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/golang-wsl .


CMD ["./golang-wsl"]

