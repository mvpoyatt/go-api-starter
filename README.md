GO STUFF:
- run$ go mod tidy

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
- Commit to Git
- Finish deleteUser function
- Determine if I want to return deltedAt as part of user
- Try to hit connct-go endpoint from go-grpc client (regenerate that code)