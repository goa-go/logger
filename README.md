# logger
A simple logging middleware for goa.

```
[2019-09-03 11:31:07] <-- GET /
[2019-09-03 11:31:07] --> GET / 200 75ms
[2019-09-03 11:31:15] <-- GET /
[2019-09-03 11:31:15] xxx GET /error 500 60ms
```

## Installation

```bash
$ go get -u github.com/goa-go/goa
```
## Notes

Recommended that you .use() this middleware near the top to "wrap" all subsequent middleware.

## Example
```go
package main

import (
  "github.com/goa-go/goa"
  "github.com/goa-go/logger"
)

func main() {
  app := goa.New()

  app.Use(logger.New())
  ...
}
```

## License

[MIT](https://github.com/goa-go/logger/blob/master/LICENSE)
