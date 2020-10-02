// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

// Config Structure
type Config struct {
	Team     string `config:"team"`
	Status   string `config:"status"`
	Pipeline string `config:"pipeline"`
	Project  string `config:"project"`
	Error    string `config:"error"`
}

// DefaultConfig Default Values
var DefaultConfig = Config{}
