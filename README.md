# TODO

* Cleanup old Firebase deployments

# Testing Locally

```bash
# Test go code locally
go run main.go
# Test Cloud Functions locally
docker run -it -v ${PWD}:/go/src/github.com/bradwh
itfield/FirebaseDeployBot/ --workdir /go/src/github.com/bradwhitfield/FirebaseDeployBot/ golang:stretch bash
go get golang.org/x/net/context
go get golang.org/x/oauth2/google
go get google.golang.org/api/cloudbuild/v1
cat request.json | cloud-functions-go-shim -entry-point F -event-type http -plugin-path functions.so
```
