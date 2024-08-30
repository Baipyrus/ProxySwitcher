package util

type Config struct {
	Name  string    `json:"name"`
	Cmd   string    `json:"cmd"`
	Set   []Variant `json:"set,omitempty"`
	Unset []Variant `json:"unset,omitempty"`
}

type Variant struct {
	Arguments []string `json:"args"`
	Type      string   `json:"type,omitempty"`
	Equator   string   `json:"equator,omitempty"`
}
