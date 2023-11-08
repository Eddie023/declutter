package config

type Flags = map[string]bool

type Config struct {
	Path  string
	Flags Flags
}
