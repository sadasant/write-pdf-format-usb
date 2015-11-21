go get github.com/rakyll/statik
cd server
rmdir statik /s /q
%GOPATH%\bin\statik.exe -src=%GOPATH%\src\github.com\sadasant\write-pdf-format-usb\public
cd ..
go build -o bin\app.exe
