go get github.com/rakyll/statik
cd server && statik -src=../public && cd ..
go build -o bin/app 
