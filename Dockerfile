FROM golang:latest as builder
RUN mkdir /app 
WORKDIR /app 
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .
# RUN go build -o main . 
# EXPOSE 8080
# CMD ["/app/main"]

FROM scratch
COPY --from=builder /main ./
ENTRYPOINT ["./main"]
