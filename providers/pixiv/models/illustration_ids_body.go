package models

import "encoding/json"

type IllustrationIDsBody struct {
	/**
	 * Using `json.RawMessage` here skips reading and processing data which saves
	 * on allocations and CPU resources. This is the most efficient solution when
	 * we expect for values to be null.
	 *
	 * - Using `map[string]struct{}` would correctly parse "null" as `nil`,
	 *   but will break if Pixiv starts sending primitives.
	 * - Using `map[string]any` would try to parse and process value using
	 *   reflection which will waste CPU cycles on things that we don't need.
	 */
	Illusts map[string]json.RawMessage `json:"illusts"`
}
