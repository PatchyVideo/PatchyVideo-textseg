go get -d -v
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o textseg.app
docker build --no-cache -t patchyvideo-textseg:latest .
