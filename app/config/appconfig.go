package config

type AppConfig struct {
	App    AppTable    `toml:"app"`
	Bot    BotTable    `toml:"bot"`
	Server ServerTable `toml:"server"`
}

type AppTable struct {
	Debug bool `toml:"debug"`
}

type BotTable struct {
	Token   string `toml:"token"`
	OnError string `toml:"on_error"`
	OnStop  string `toml:"on_stop"`
}

type ServerTable struct {
	Host string `toml:"host"`
	Port uint16 `toml:"port"`
}
