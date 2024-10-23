$env:GOARCH="amd64"; $env:GOOS="linux"; go build -ldflags="-s -w" -o bootstrap main.go; Compress-Archive -Path bootstrap -DestinationPath main.zip

