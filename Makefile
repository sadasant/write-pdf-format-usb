all: statik app

statik:
	go get github.com/rakyll/statik
	cd server && statik -src=../public

app:
	go build -o bin/app 
