# Portable URL Shortener

This project has the objective of learning how to build, document, and test web APIs written in Golang.

* [ ] Functional
	* [X] Accept URL to short from a POST request and provvide the shortened url.
	* [X] When using that url, it will redirect to the appropriate url.
	* [ ] Show url statistics (time opened, time remaining) via web interface and API endpoint.
* [X] Non-functional
	* [X] API
		* [X] Version the API
		* [X] Document API endpoints
	* [X] Deployment
	  * [X] Create configurable base form url (gs.x/uniqueid).
	  * [X] Create public containerized envirioment to deploy not only in a more pratical way but also on personal home lab.

## Folder structure
- ***Api***: Will be the folder in which all the api endpoints, documentation and relative tests will be.
- ***Pkg***: Will be the folder in which all the common code for the application is stored.
- ***Cmd***: Will be the folder in which we will find al excecutables necessary in order to start the shortener.
- ***Website***: Will be the folder in which the static html files will be placed.

## Building Instruction

In alternative it is possible to build the binary
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X 'shortic/pkg/deployment.version="0.0.1"' -X 'shortic/pkg/deployment.commitHash=$(git rev-parse HEAD)' -X 'shortic/pkg/deployment.buildUser=$(git config user.name)' -X 'shortic/pkg/deployment.buildTime=$(date)'" -o ./build/shortic ./cmd/restapi/main.go
```

### Self Host Docker Container

In order to build the application this is the tested command:
 ```
 sudo docker build \
 --build-arg version=0.0.1 \
 --build-arg commitHash="$(git rev-parse HEAD)" \
 --build-arg buildUser="$(git config user.name)" \``
  --build-arg buildTime="$(date)" \
 -t Shortic \
  .
```
Or use the public image (Not implemented)
