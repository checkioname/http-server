package core

import (
	"errors"
	"net/http"
)

type Plugin interface {
	Init(config map[string]interface{}) error
	Name() string
}

type PluginHTTP interface {
	Plugin
	Name() string
	Match(r *http.Request) bool
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

var registry = map[string]PluginHTTP{}

func RegisterModule(plugin PluginHTTP) {
	name := plugin.Name()
	if _, exist := registry[name]; exist {
		panic("duplicate plugin name: " + name)
	}
	registry[plugin.Name()] = plugin
}

func GetPlugin(name string) (PluginHTTP, error) {
	plugin, exist := registry[name]
	if !exist {
		return nil, errors.New("plugin not found: " + name)
	}
	return plugin, nil
}

func GetAllPlugins() []PluginHTTP {
	var list []PluginHTTP
	for _, plugin := range registry {
		list = append(list, plugin)
	}
	return list
}
