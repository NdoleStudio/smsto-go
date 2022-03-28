# smsto-go

[![Build](https://github.com/NdoleStudio/smsto-go/actions/workflows/main.yml/badge.svg)](https://github.com/NdoleStudio/smsto-go/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/NdoleStudio/smsto-go/branch/main/graph/badge.svg)](https://codecov.io/gh/NdoleStudio/smsto-go)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/NdoleStudio/smsto-go/badges/quality-score.png?b=main)](https://scrutinizer-ci.com/g/NdoleStudio/smsto-go/?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/NdoleStudio/smsto-go)](https://goreportcard.com/report/github.com/NdoleStudio/smsto-go)
[![GitHub contributors](https://img.shields.io/github/contributors/NdoleStudio/smsto-go)](https://github.com/NdoleStudio/smsto-go/graphs/contributors)
[![GitHub license](https://img.shields.io/github/license/NdoleStudio/smsto-go?color=brightgreen)](https://github.com/NdoleStudio/smsto-go/blob/master/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/NdoleStudio/smsto-go)](https://pkg.go.dev/github.com/NdoleStudio/smsto-go)


This package provides a `go` client for the SMS.to HTTP API https://developers.sms.to

## Installation

`smsto-go` is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/NdoleStudio/smsto-go
```

Alternatively the same can be achieved if you use `import` in a package:

```go
import "github.com/NdoleStudio/smsto-go"
```


## Implemented

- [SMS](#sms)
    - `POST /sms/send`: Send single message to a number

## Usage

### Initializing the Client

An instance of the client can be created using `New()`.

```go
package main

import (
	"github.com/NdoleStudio/smsto-go"
)

func main()  {
	client := smsto.New(smsto.WithAPIKey(/* API KEY */))
}
```

### Error handling

All API calls return an `error` as the last return object. All successful calls will return a `nil` error.

```go
result, response, err := client.SMS.SendSingle(context.Background(), &smsto.SmsSendSingleRequest{})
if err != nil {
    //handle error
}
```

### SMS

#### `POST /sms/send`: Send single message to a number

```go
status, response, err := client.SMS.SendSingle(context.Background(), &SmsSendSingleRequest{
    Message:      "This is test and \n this is a new line",
    To:           "+35799999999999",
})

if err != nil {
	log.Fatal(err)
}

log.Println(status.Success) // true
```

## Testing

You can run the unit tests for this client from the root directory using the command below:

```bash
go test -v
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
