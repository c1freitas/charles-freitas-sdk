# charles-freitas-sdk


This is a GO SDK that can be used for accessing the API at https://the-one-api.dev/. Refer to the [Documentation](https://the-one-api.dev/documentation) for specific details on the API.

## Installation

You can use go to install this in your GOPATH or as a module in your existing application.

```
 go get github.com/c1freitas/charles-freitas-sdk
```


## Usage

Most of the endpoints will require an API Access token which can be obtained by creating an account at https://the-one-api.dev/login

Example of creating and using the API can be found in the ./client/client_test.go module. 

Basic usage examples:

```
    // create new client
	client := client.NewClient(os.Getenv("ACCESS_TOKEN"))
	bookData, err := client.GetBooks(nil)
```
All active endpoints have basic unit tests and are a good refence for how to use the client package.


## Current State

This is currently a work in progress with most of, but not all, of the endpoints oprational. Sorting and Pagination 
are operational, but more work could be done around optimizing the Client and supporting additional 
configuration (ie timeout, retry logic, etc)


