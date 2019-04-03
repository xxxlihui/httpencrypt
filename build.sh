export GOOS=linux
go build -o httpencrypt
export GOOS=windows
go build -o httpencrypt.exe