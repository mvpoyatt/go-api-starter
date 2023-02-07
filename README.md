
# Go API Starter

This library is a starter for an API implemented in GoLang. It comes with user login functionality, logging capablities, environment-based configurations. If login functionality isn't required it can easily be used as an example and replaced with appropriate endpoint using the steps outlined below. Stack includes grpc, buf/connect-go, MySQL, and GORM. There is also an accompanying front-end starter built using Qwik to demo a typescript/client implementation of connect-go: [qwik-client-starter](https://github.com/mvpoyatt/qwik-client-starter) (WIP)

## Run As-Is

**1. Install Go and Dependencies**
- [Install Go](https://go.dev/doc/install)
- Clone this repo, cd into directory and run
	```bash
	$ go get github.com/mvpoyatt/go-api/api
	$ go get github.com/mvpoyatt/go-api/configs
	$ go get github.com/mvpoyatt/go-api/database
	$ go mod tidy
	```
	
**2. Download MySQL and MySQL Workbench**
- [MySQL](https://dev.mysql.com/downloads/mysql/) - Replace ```db:password``` in ```/configs/development.yaml``` with password you enter here
- [MySQL Workbench](https://dev.mysql.com/downloads/workbench/) - Create new schema named ```testdb``` or replace ```db:dbname``` in ```/configs/development.yaml``` with name of new schema

**3. Start Server**
- In root directory of project run
	```bash
	$ go run entry/main.go
	```

## To Work with Buf

1. [Install buf tools](https://docs.buf.build/installation)
	- (for Homebrew) ```$ brew install bufbuild/buf/buf```
2. [Install connect tools](https://connect.build/docs/go/getting-started/)
	- (for mac)
	```bash
	$ go install github.com/bufbuild/buf/cmd/buf@latest
	$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    $ go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
    $ [ -n "$(go env GOBIN)" ] && export PATH="$(go env GOBIN):${PATH}"
    $ [ -n "$(go env GOPATH)" ] && export PATH="$(go env GOPATH)/bin:${PATH}"
	```
3. Remove existing code and re-generate
	```bash
	$ rm -rf gen/
	$ buf generate
	```
	- This generates code in ```gen/``` folder. Delete and re-run after changing ```.proto``` file.

## Creating New Endpoints


1. Add new API folders to buf.work.yaml
2. Update buf.gen.yaml for new:
	  - Code generating plugins
	  - External packages (exlude default go_option prefixes)
3. In home directory of project run
	```bash
	rm -rf gen  # reset existing generated code
	$ buf lint
	$ buf build
	$ buf generate
	```

## To-Do
- Make DB conenction SSL once on server
- Revisit API configs when deploying
- Consider handling user auth in interceptor
- Move JWT expiry time, secret-key and password length to config file


## License

[MIT](https://choosealicense.com/licenses/mit/)