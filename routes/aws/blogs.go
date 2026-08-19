package aws

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imcrazytwkr/feedhub/models"
	"github.com/imcrazytwkr/feedhub/utils/ginutil"
)

func (r *AWSRouter) handleBlogs(c *gin.Context) {
	category, lang := resolveBlogsPath(
		strings.TrimRight(c.Param("slug"), "/"),
		strings.TrimRight(c.Param("lang"), "/"),
	)

	feed, err := r.provider.GetBlogs(c.Request.Context(), category, lang)
	if err != nil {
		ginutil.HandleError(c, err)
		return
	}

	if feed == nil {
		c.String(http.StatusNotFound, "AWS: no blog posts found")
		return
	}

	ginutil.RenderFeed(c, feed)
}

// These cases are excessively documented because nested optional params are
// terrible from code standpoint but good for UX
func resolveBlogsPath(category string, lang string) (string, models.Language) {
	// `//:lang` is normalized to `/:lang` by GIN so it is impossible for "lang"
	// to be set when "category" is empty.
	if len(category) == 0 {
		// No parameters (all categories in English)
		return "", models.LanguageEn
	}

	if len(lang) == 0 {
		language := models.ParseLanguage(category)

		// `/aws/:lang` (all posts in language)
		if language != models.LanguageUnknown {
			return "", language
		}

		// `/aws/:category` (posts for category in English)
		return category, models.LanguageEn
	}

	// `/aws/:category/:language` (posts for category in language)
	language := models.ParseLanguage(lang)
	if language == models.LanguageUnknown {
		return category, models.LanguageEn
	}

	return category, language
}
