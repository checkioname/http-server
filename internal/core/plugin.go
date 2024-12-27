package core


// every plugin should implement this interface
type Plugin interface {
  RegisterModule() error
  Init(config map[string]interface{}) error
  Name() string
}
