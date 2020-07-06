package translation

import (
	"github.com/gin-gonic/gin"
	"github.com/spacetab-io/i18n-go/translation"
)

const (
	ContextParam = "translation-context"
)

func Header(base translation.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ContextParam, translation.NewContext(base, c.Request))
	}
}

func ContextFromGin(base translation.Context, c *gin.Context) translation.Context {
	v, ok := c.Get(ContextParam)
	if !ok {
		return base
	}
	result, ok := v.(translation.Context)
	if !ok {
		return base
	}

	return result
}
