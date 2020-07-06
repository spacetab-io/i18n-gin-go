# i18n-gin-go
A gin wrappers for a i18n-go package

# Translation

## Usage

Basic usage a gin middleware for processing HTTP's requests

```go
package pack

import (
	"encoding/json"
	"net/http"

 	"github.com/gin-gonic/gin"
	translation_gin "github.com/spacetab-io/i18n-gin-go/translation"
	"github.com/spacetab-io/i18n-go/translation"
)

var (
	// a configuration of default properties of translation context
	// it could be loaded from json or yaml configuration file
	Conf = translation.Conf{
		Display: "ru",
		Fallback: "en",
		Second:  "en",
		TranslationList: false,
	}
)

// A model definition

type Record struct {
	Id   int                `json:"id"`                  
	Name translation.String `json:"name"`
}

func (o *Record) ApplyTranslationCtx(ctx translation.Context) {
	o.Name.ApplyTranslationCtx(ctx)
}

// A router definition

func Router() http.Handler {
	router := gin.New()

	router.Use(translation_gin.Header(&Conf))

	router.GET("/record", Handler)
	
	return router
}

// A handler of an http request

func Handler(c *gin.Context) {
	rec := &Record{
	    Id: 10,
		Name: translation.String {
		    Translate: map[string]string {
    			"en": "Hello world!",
    			"ru": "Здравствуй мир!",
	    	},
	    },
	}
	
	rec.ApplyTranslationCtx(translation_gin.ContextFromGin(&Conf, c))
	
	c.JSON(http.StatusOK, rec)
}
```
