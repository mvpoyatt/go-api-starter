After cloning repo...

1. Install Go
  a. Download go: https://go.dev/doc/install
  b. Install dependencies:
    cd into project directory and run
    $ go get github.com/mvpoyatt/go-api/api
    $ go get github.com/mvpoyatt/go-api/configs
    $ go mod tidy
    $ go run entry/main.go

2. To work with buf
  a. Install buf tools: https://docs.buf.build/installation
    (for mac): 
    $ brew install bufbuild/buf/buf
  b. Install connect tools: https://connect.build/docs/go/getting-started/
    (for mac):
    $ go install github.com/bufbuild/buf/cmd/buf@latest
    $ go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    $ go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
    $ [ -n "$(go env GOBIN)" ] && export PATH="$(go env GOBIN):${PATH}"
    $ [ -n "$(go env GOPATH)" ] && export PATH="$(go env GOPATH)/bin:${PATH}"
  c. Generate helper code:
    $ buf generate 
    Generates code in gen/ folder. Delete and re-run after changing .proto file

3. Download MySQL and MySQL Workbench
  a. https://dev.mysql.com/downloads/mysql/
    Replace db:password in /configs/development.yaml with password entered here
  b. https://dev.mysql.com/downloads/workbench/
    Create new schema named testdb, or replace db:dbname in /configs/development.yaml with name of new schema.

4. Start server
  $ go run entry/main.go

BUF STUFF:
- Add new API folders to buf.work.yaml
- Update buf.gen.yaml for new:
  - Code generating plugins
  - External packages (exlude default go_option prefixes)
- run$ buf build
- run$ buf generate
- run$ rm -rf gen // to reset generated code

TODO:
- Make DB connection SSL once on server
- Publish final API to BRS

DO NEXT:
- Finish deleteUser function
- Determine if I want to return deltedAt as part of user
- Try to hit connct-go endpoint from go-grpc client (regenerate that code)