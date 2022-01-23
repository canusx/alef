package alef

// Config is the configuration for the alef plugin.
type Config struct {
	// add your config fields
	mongoURI string
}

// Validate validates the configuration, and return an error if it is invalid.
func (c *Config) Validate() error {
	return c.Validate()
}

// DefaultConfig is the default configuration.
var DefaultConfig = Config{}

func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return c.Validate()
}
