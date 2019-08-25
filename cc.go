package config_center

type ConfigReader interface {
	Init() error
	Read(key string) ([]*ConfigValue, error)
	Close() error
}

type ConfigValue struct {
	Raw    []byte
	Key    string
	Prefix string
	RawKey string
}
