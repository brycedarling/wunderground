# go-wunderground-client

API Client for the Wunderground API ten day forecast and astronomy endpoints.

You must set a WUNDERGROUND_API_KEY environment variable to use this program.

Example Usage:

* wunderground tendayforecast "Asheville, NC"
* wunderground astronomy "Bozeman, MT"

Build and Test Instructions:

* In order to build, you will need to set your GOPATH and run `go build`
* In order to run the tests, you will also need GOPATH set and run `go test`

#### TODO

* The ten day forecast returns a massive JSON object, the entire thing is not currently being json decoded.
* The ten day forecast and astronomy Get*Data methods could probably have a little better tests.
* Fix the weirdness of using main and wunderground packages and the extra directories, related to GOPATH.
* Add concurrency with goroutines/channels

