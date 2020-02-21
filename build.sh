go get -d -v
go build -o textseg.app
docker build --no-cache -t patchyvideo-textseg:latest .
