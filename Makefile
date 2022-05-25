all:
	go build 

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build

clean:
	rm -f ./main
	rm -f ./tinfo-go
	rm -rf ./Tinfo-go.release.zip
	rm -rf ./releases/*
	ls -lah

release:
	zip -r ./Tinfo-go.release.zip ./tinfo-go ./site_conf.json
	unzip ./Tinfo-go.release.zip -d ./releases/
	ls -lah