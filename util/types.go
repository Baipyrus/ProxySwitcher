package util

type Config struct {
	Name  string     `json:"name"`
	Cmd   string     `json:"cmd"`
	Set   []*Variant `json:"set,omitempty"`
	Unset []*Variant `json:"unset,omitempty"`
}

type Type string

const (
	TEXT     Type = "text"
	VARIABLE Type = "variable"
)

type Variant struct {
	Arguments []string `json:"args"`
	Type      Type     `json:"type,omitempty"`
	Equator   string   `json:"equator,omitempty"`
}

type Command struct {
	Name      string
	Arguments []string
}
