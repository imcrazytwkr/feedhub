package next_preflight

type InlineFlightTag int

const (
	InlineFlightTagBootstrap = iota
	InlineFlightTagData
	InlineFlightTagFormState
	InlineFlightTagBinary
)
