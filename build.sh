export GOOS=linux
go build -o httpencrypt
echo 编译linux版本成功 httpencrypt
export GOOS=windows
go build -o httpencrypt.exe
echo 编译windows版本成功 httpencrypt.exe
