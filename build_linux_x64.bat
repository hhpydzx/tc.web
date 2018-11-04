set GOARCH=amd64
set GOOS=linux
go build -o tc.web main.go
echo "编译完成，任意键退出"
pause