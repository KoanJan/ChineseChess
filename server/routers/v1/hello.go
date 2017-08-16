package v1

import "github.com/kataras/iris/context"

// Hello echoes "hello, iris!"
func Hello(ctx context.Context) {

	ctx.WriteString("hello, iris!")
}
