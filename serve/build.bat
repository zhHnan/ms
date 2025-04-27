chcp 65001
@echo off
:loop
@echo off&amp;color 0A
cls
echo,
echo 请选择要编译的系统环境：
echo,
echo 1. Windows_amd64
echo 2. linux_amd64

set/p action=请选择:
if %action% == 1 goto build_Windows_amd64
if %action% == 2 goto build_linux_amd64

:build_Windows_amd64
echo 编译Windows版本64位
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -o ms-user/target/ms-user.exe ms-user/main.go
go build -o ms-api/target/ms-api.exe ms-api/main.go
go build -o ms-api/target/ms-project.exe ms-project/main.go
:build_linux_amd64
echo 编译Linux版本64位
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o ms-user/target/ms-user ms-user/main.go
go build -o ms-api/target/ms-api ms-api/main.go
go build -o ms-api/target/ms-project ms-project/main.go