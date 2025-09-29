package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

func SanitizerMiddleware() gin.HandlerFunc {
	policy := bluemonday.UGCPolicy()
	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodPost ||
			ctx.Request.Method == http.MethodPut ||
			ctx.Request.Method == http.MethodPatch {
			if err := ctx.Request.ParseForm(); err == nil {
				for k, vals := range ctx.Request.PostForm {
					cleanVals := make([]string, len(vals))
					for i, v := range cleanVals {
						cleanVals[i] = policy.Sanitize(v)
					}
					ctx.Request.PostForm[k] = cleanVals
				}
			}
		}
		ctx.Next()
	}
}
