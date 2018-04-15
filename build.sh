#/bin/sh
set -e

env GOOS=linux GOARCH=386 go build auth.go utils.go
env GOOS=linux GOARCH=386 go build clean.go utils.go