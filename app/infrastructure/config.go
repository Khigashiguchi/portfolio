package infrastructure

import (
    "github.com/spf13/viper"
    "fmt"
)

// DBConfig represents configuration information.
type DBConfig struct {
    Host string
    Name string
    User string
    Password string
    Port int
    SQLMode string
}

// Config represents global configuration.
type Config struct {
    DB DBConfig
}

// NewConfig return Config struct.
func NewConfig(appEnv string, confDir string) (Config, error) {
    conf := Config{}

    viper.SetConfigName(appEnv)
    viper.AddConfigPath(confDir)
    err := viper.ReadInConfig()
    if err != nil {
        return conf, fmt.Errorf("failed to load configuration file in %s dir, given appEnv: %s, given err: %s", confDir, appEnv, err)
    }

    conf.DB = DBConfig{
        Host: viper.GetString("DB.Host"),
        Port: viper.GetInt("DB.Port"),
        Name: viper.GetString("DB.Name"),
        User: viper.GetString("DB.User"),
        Password: viper.GetString("DB.Password"),
        SQLMode: viper.GetString("DB.SQLMode"),
    }

    return conf, nil
}
