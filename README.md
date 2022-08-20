# Portable URL Shortener

This project has the objective of learning how to build, document, and test web APIs written in Golang.

* [ ] Functional
	* [X] Accept URL to short from a POST request and provvide the shortened url.
	* [X] When using that url, it will redirect to the appropriate url.
	* [ ] Show url statistics (time opened, time remaining) via web interface and API endpoint.
* [ ] Non-functional
	* [ ] API
		* [X] Version the API
		* [ ] Implemente API Health checks for each service
		* [X] Document API endpoints
	* [ ] Testing
	  * [ ] Test basic edge cases of the backend.
	  * [ ] Test possible user inputs throught form.
	  * [ ] Test API endpoints.
	* [ ] Deployment
	  * [ ] Create configurable redirect url ex.(goorter.io).
	  * [ ] Create configurable base form url (gs.x/uniqueid).
	  * [ ] Create public containerized envirioment to deploy not only in a more pratical way but also on personal home lab.
	  * [ ] Create testing and building pipeline (optional).

## Folder structure
- ***Api***: Will be the folder in which all the api endpoints, documentation and relative tests will be.
- ***Pkg***: Will be the folder in which all the common code for the application is stored.
- ***Cmd***: Will be the folder in which we will find al excecutables necessary in order to start the shortener.
- ***Website***: Will be the folder in which the static html files will be placed.

