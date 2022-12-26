package config

type DB struct {
	File string `config:"file" json:"file"`
}

type Config struct {
	Port int `config:"port" json:"port"`
	DB   DB  `config:"db" json:"db"`
}
