package next_preflight

type StreamParser struct {
	Rows  map[int]*Row
	Order []int
}
