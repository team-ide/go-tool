@echo on

cd %~dp0
cd ../

go test -v -run Test* ./datamove
