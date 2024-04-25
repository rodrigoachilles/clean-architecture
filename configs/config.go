package configs

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Conf struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort    string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort string `mapstructure:"GRAPHQL_SERVER_PORT"`
	MQHost            string `mapstructure:"MQ_HOST"`
	MQPort            string `mapstructure:"MQ_PORT"`
	MQUser            string `mapstructure:"MQ_USER"`
	MQPassword        string `mapstructure:"MQ_PASSWORD"`
}

func LoadConfig() (*Conf, error) {
	targetFileName := ".env"
	var cfg *Conf

	exeDir, _ := os.Executable()
	rootDir := filepath.Dir(exeDir)

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(rootDir)
	viper.SetConfigFile(targetFileName)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, err
}
