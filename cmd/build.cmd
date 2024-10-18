@echo on

cd %~dp0
cd ../


:: -s 去掉符号表
:: -w 去掉调试信息
:: -v 编译时显示包名
:: -p n 开启并发编译，默认情况，n为CPU逻辑核数
:: -a 强制重新构建
:: -n 打印编译时会用到的所有命令，但不真正执行


go build -o go-tool.exe .

go build -ldflags="-s" -o go-tool-s.exe .
go build -ldflags="-w" -o go-tool-w.exe .
go build -ldflags="-w -s" -o go-tool-s-w.exe .
