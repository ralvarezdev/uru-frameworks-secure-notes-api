@echo off
echo "Compiling server..."
cd ./cmd/server
go build -o ../../bin/server
cd ../../
echo "Compiling server... Done"