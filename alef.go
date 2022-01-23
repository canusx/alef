package alef

import (
	"github.com/DrmagicE/gmqtt/config"
	"github.com/DrmagicE/gmqtt/server"
	"go.uber.org/zap"
)

var _ server.Plugin = (*Alef)(nil)

const Name = "alef"

func init() {
	server.RegisterPlugin(Name, New)
	config.RegisterDefaultPluginConfig(Name, &DefaultConfig)
}

func New(config config.Config) (server.Plugin, error) {
	a := &Alef{}
	return a, nil
}

var log *zap.Logger

func (a *Alef) Load(service server.Server) error {
	log = server.LoggerWithField(zap.String("plugin", Name))
	return nil
}

func (a *Alef) Unload() error {
	return nil
}

func (a *Alef) Name() string {
	return Name
}

type Alef struct {
	config *Config
}
