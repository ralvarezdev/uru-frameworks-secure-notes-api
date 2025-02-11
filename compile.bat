@echo off
echo "Compiling server..."
cd ./cmd/server
go mod tidy
go build -o ../../bin/server
cd ../../
echo "Compiling server... Done"