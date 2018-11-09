#!/bin/bash
echo compiling ...
env GOOS=linux GOARCH=arm GOARM=6 go build -o main main.go

echo uploading ...
scp main 172.16.0.1:~/P4wnP1/build

