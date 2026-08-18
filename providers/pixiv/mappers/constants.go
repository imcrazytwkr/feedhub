package mappers

import "regexp"

var imageRe = regexp.MustCompile(`(/\d{4}/(?:\d{2}/){5}\d+)_p\d+[^\.]*\.(\w{3,})$`)

const artistPrefix = "https://www.pixiv.net/en/users/"
const postPrefix = "https://www.pixiv.net/en/artworks/"
const cdnPrefix = "https://i.pixiv.cat"
