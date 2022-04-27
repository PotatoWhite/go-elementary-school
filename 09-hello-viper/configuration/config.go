package configuration

var RuntimeConf = Config{}

type Config struct {
	Datasource Datasource `yaml:"datasource"`
	Server     Server     `yaml:"server"`
}

type Datasource struct {
	Dsn string `yaml:"dsn"`
}

type Server struct {
	Port int `yaml:"port"`
}
