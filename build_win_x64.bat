set GOARCH=amd64
set GOOS=windows
go build -o tc.web.exe main.go
echo "编译完成，任意键退出"
pause