// Package florest-core/src provides a workflow based REST API framework.
//
// Pre-requisites
//
//
// 1. Go 1.5+
//
// 2. Linux or MacOS
//
// 3. https://onsi.github.io/ginkgo/ for executing the tests
//
// Usage
//
// Clone the repo:-
//  	cd <GOPROJECTSPATH>
//  	git clone https://github.com/jabong/florest-core
//
// Bootstrap the new application to be created from florest Let's assume
// "APPDIR" is the absolute location where new application's code will reside.
// For example, let's say the new application to be created is named restapi to
// be placed in "/Users/tuk/golang/src/github.com/jabong/" then "APPDIR" denotes
// the location "/Users/tuk/golang/src/github.com/jabong/restapi"
//  	cd <GOPROJECTSPATH>/florest-core
//  	make newapp NEWAPP="APPDIR"
// The above will create a new app based on "florest" with the necessary structure.
//
// Set-up the application log directory. Let's say if the application was created
// as mentioned in the previous step then this will look like below:-
//  	sudo mkdir /var/log/restapi/          # This can be changed
//  	chown <user who will be executing the app> /var/log/restapi
//
// To build the application execute the below command:-
//  	cd APPDIR
//  	make deploy
// If the above command is successful then this will create a binary named after the application
// name. In this case the binary will be named as "restapi". The binary can be executed using the
// below command:-
//
//  	cd APPDIR/bin
//  	./restapi
//
// Tests
//
// To run the tests execute the below command:-
//		cd APPDIR
//		make test
//
// To get the test coverage execute the below command:-
//
//		cd APPDIR
//		make coverage
//
// To execute all the benchmark tests:-
//		make bench
//
// The application code can also be formatted as per gofmt. To do this execute the below command:-
//
//		make format
//
// Refer https://github.com/jabong/florest-core/wiki for more
package src
