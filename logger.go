package logger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/goa-go/goa"
)

var (
	bold   = color.New(color.Bold).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
)

func New() goa.Middleware {
	return func(c *goa.Context, next func()) {
		start := time.Now()
		fmt.Fprintf(
			color.Output,
			"[%s] <-- %s %s\n",
			start.Format("2006-01-02 15:04:05"),
			bold(c.Method),
			c.URL,
		)
		defer func() {
			if err := recover(); err != nil {
				statusCode := 500
				if e, ok := err.(goa.Error); ok {
					statusCode = e.Code
				}
				fmt.Fprintf(
					color.Output,
					"[%s] %s %s %s %s %d%s\n",
					time.Now().Format("2006-01-02 15:04:05"),
					red("xxx"),
					bold(c.Method),
					c.Path,
					ColorStatus(statusCode),
					time.Since(start).Nanoseconds()/1e6,
					"ms",
				)
				panic(err)
			}
		}()
		next()
		fmt.Fprintf(
			color.Output,
			"[%s] %s %s %s %s %d%s\n",
			time.Now().Format("2006-01-02 15:04:05"),
			"-->",
			bold(c.Method),
			c.Path,
			ColorStatus(c.GetStatus()),
			time.Since(start).Nanoseconds()/1e6,
			"ms",
		)
	}
}

func ColorStatus(code int) string {
	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return green(code)
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return cyan(code)
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return yellow(code)
	default:
		return red(code)
	}
}
