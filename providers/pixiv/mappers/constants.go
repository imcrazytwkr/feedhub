package mappers

import (
	"regexp"
)

const errorKey = "error"
const messageKey = "message"

const bodyKey = "body"
const canonicalKey = "canonical"
const createDateKey = "createDate"
const updateDateKey = "updateDate"
const descriptionKey = "description"
const descriptionHeaderKey = "descriptionHeader"
const extraDataKey = "extraData"
const illustsKey = "illusts"
const metaKey = "meta"
const pageCountKey = "pageCount"
const titleKey = "title"
const urlKey = "url"
const userIdKey = "userId"
const userNameKey = "userName"
const worksKey = "works"

var imageRe = regexp.MustCompile(`(/\d{4}/(?:\d{2}/){5}\d+)_p\d+[^\.]*\.(\w{3,})$`)

const artistPrefix = "https://www.pixiv.net/en/users/"
const postPrefix = "https://www.pixiv.net/en/artworks/"
const cdnPrefix = "https://i.pixiv.cat"
