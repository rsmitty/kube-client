all: darwin linux

darwin:
	env GOOS=darwin GOARCH=amd64 go build -o bin/kube-client-darwin .

linux:
	env GOOS=linux GOARCH=amd64 go build -o bin/kube-client-linux .

clean:
	rm -rf ./bin