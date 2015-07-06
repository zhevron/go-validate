go-validate - Object validation library
=======================================

[![wercker status](https://app.wercker.com/status/75af4a3e778a36be820d46a947868b89/s "wercker status")](https://app.wercker.com/project/bykey/75af4a3e778a36be820d46a947868b89)
[![Coverage Status](https://coveralls.io/repos/zhevron/go-validate/badge.svg?branch=HEAD)](https://coveralls.io/r/zhevron/go-validate)
[![GoDoc](https://godoc.org/gopkg.in/zhevron/go-validate.v0/validate?status.svg)](https://godoc.org/gopkg.in/zhevron/go-validate.v0/validate)

**go-validate** is an object validation library for [Go](https://golang.org/).  

For package documentation, refer to the GoDoc badge above.

## Installation

```
go get gopkg.in/zhevron/go-validate.v0/validate
```

## Usage

```go
package main

import (
  "fmt"

  "gopkg.in/zhevron/go-validate.v0/validate"
)

type MyObject struct {
  String string `minLen:"10" maxLen:"50"`
  Int    int    `min:"5"`
  Error  error  `nil:"false"`
}

func main() {
  var obj MyObject

  if ok, errs := validate.Validate(obj); !ok {
    fmt.Println("The following validation errors occured:")

    for _, err := range errs {
      fmt.Printf("  %s\n", err.Error())
    }
  }
}
```

## License

**go-validate** is licensed under the [MIT license](http://opensource.org/licenses/MIT).
