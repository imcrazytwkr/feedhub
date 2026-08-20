package preflight

import (
	"github.com/imcrazytwkr/feedhub/utils/jsontree"
	np "github.com/imcrazytwkr/feedhub/utils/next_preflight"
)

type preflightArticle struct {
	Type           string
	Title          string
	Body           *jsontree.Value
	FootnotesTitle string
	FootnotesBody  *jsontree.Value
}

func (a *preflightArticle) CanRender() bool {
	return a != nil && hasBlocks(a.Body) && (a.Type == "post" || a.Type == "engineeringArticle")
}

func hasBlocks(v *jsontree.Value) bool {
	switch v.Type() {
	case jsontree.TypeArray:
		return len(v.GetArray()) > 0
	case jsontree.TypeObject:
		return v.GetObject().Len() > 0
	default:
		return false
	}
}

func findPreflightArticle(preflight []byte) *preflightArticle {
	stream, err := np.Parse(preflight)
	if err != nil || stream == nil {
		return nil
	}

	for _, id := range stream.Order {
		row := stream.Rows[id]
		if row == nil || row.Tag != 0 {
			continue
		}

		for v := range walkJSON(stream.Resolve(id)) {
			if article := asPreflightArticle(v); article.CanRender() {
				return article
			}
		}
	}

	return nil
}

func asPreflightArticle(v *jsontree.Value) *preflightArticle {
	if v.GetObject().Len() == 0 {
		return nil
	}

	if nested := v.Get("post"); nested.GetObject().Len() > 0 {
		if article := decodeArticle(nested); article.CanRender() {
			return article
		}
	}

	if nested := v.Get("article"); nested.GetObject().Len() > 0 {
		if article := decodeArticle(nested); article.CanRender() {
			return article
		}
	}

	return decodeArticle(v)
}

func decodeArticle(v *jsontree.Value) *preflightArticle {
	if v.GetObject().Len() == 0 {
		return nil
	}

	return &preflightArticle{
		Type:           v.GetString("_type"),
		Title:          v.GetString("title"),
		Body:           v.Get("body"),
		FootnotesTitle: v.GetString("footnotesTitle"),
		FootnotesBody:  v.Get("footnotesBody"),
	}
}
