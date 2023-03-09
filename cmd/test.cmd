@echo off

cd %~dp0
cd ../


go test -v -cover -coverprofile="cover.out" -covermode=count ./...

go tool cover -func="cover.out"

go tool cover -func="cover.out" > "cover-func.out"