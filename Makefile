all: http

http: http.go
	go build http.go

clean:
	rm http
	go clean
