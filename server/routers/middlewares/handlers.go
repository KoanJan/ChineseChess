package middlewares

import "github.com/kataras/iris/context"

const JwtTokenHttpHeaderName = "token"

var Handlers []context.Handler = []context.Handler{

	jwtHandler,
}

// jwtHandler is a middleware handler which validate jwt-token
func jwtHandler(ctx context.Context) {

	token := ctx.GetHeader(JwtTokenHttpHeaderName)
	if err := ValidateToken(token); err != nil {
		ctx.WriteString("未登录或登录已过期")
		ctx.StopExecution()
	}
}
