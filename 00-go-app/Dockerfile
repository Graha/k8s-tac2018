FROM golang:1.10
WORKDIR /src/
COPY web.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o web .

FROM alpine:latest 
WORKDIR /root/
COPY --from=0 /src/web .
EXPOSE 8080
CMD ["./web"]