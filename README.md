# Required
Required is used to validate if the field value is empty

### Install

```bash
go get github.com/tiaguinho/required
```

### Example

Simplest way to use the package:

```go
package main

import (
    "github.com/tiaguinho/required"
    "log"
)

type Test struct {
    FirstName string `json:"first_name" required:"-"`
    LastName  string `json:"last_name" required:"last name is required"`
}

func main()  {
    t := Test{
        FirstName: "Go",
        LastName: "Required",
    }
    
    if err := required.Validate(t); err != nil {
        log.Println(err)
    }
}
```

If you like to get all the validation messages with field's name to return in some API, just change to this:

```go
func main() {
    t := Test{
         FirstName: "Go",
         LastName: "Required",
     }
     
     if err, msg := required.ValidateWithMessage(t); err != nil {
         log.Println(err, msg)
     }
}
```
