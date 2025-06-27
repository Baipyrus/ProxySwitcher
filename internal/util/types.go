package util

type Config struct {
	Name  string     `json:"name"`
	Cmd   string     `json:"cmd"`
	Set   []*Variant `json:"set,omitempty"`
	Unset []*Variant `json:"unset,omitempty"`
}

type VariantType string

const (
	TEXT     VariantType = "text"
	VARIABLE VariantType = "variable"
)

type Variant struct {
	Arguments    []string    `json:"args"`
	Type         VariantType `json:"type,omitempty"`
	Equator      string      `json:"equator,omitempty"`
	Surround     string      `json:"surround,omitempty"`
	DiscardProxy bool        `json:"discard,omitempty"`
}

type Command struct {
	Name      string
	Arguments []string
}
