default: build
alias fmt := format

format:
  go fmt ./...

build: format
  @mkdir -p bin
  go build -o bin/tcs-notes

clean:
  rm -r bin
