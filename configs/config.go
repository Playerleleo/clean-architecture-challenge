package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort    string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort string `mapstructure:"GRAPHQL_SERVER_PORT"`
	RabbitMQHost      string `mapstructure:"RABBITMQ_HOST"`
	RabbitMQPort      string `mapstructure:"RABBITMQ_PORT"`
	RabbitMQUser      string `mapstructure:"RABBITMQ_USER"`
	RabbitMQPass      string `mapstructure:"RABBITMQ_PASS"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	// Tentar ler o arquivo de configuração
	err := viper.ReadInConfig()
	if err != nil {
		// Se não encontrar o arquivo, usar as variáveis de ambiente
		viper.SetDefault("DB_DRIVER", "mysql")
		viper.SetDefault("DB_HOST", "mysql")
		viper.SetDefault("DB_PORT", "3306")
		viper.SetDefault("DB_USER", "root")
		viper.SetDefault("DB_PASSWORD", "root")
		viper.SetDefault("DB_NAME", "orders")
		viper.SetDefault("WEB_SERVER_PORT", "8000")
		viper.SetDefault("GRPC_SERVER_PORT", "50051")
		viper.SetDefault("GRAPHQL_SERVER_PORT", "8080")
		viper.SetDefault("RABBITMQ_HOST", "rabbitmq")
		viper.SetDefault("RABBITMQ_PORT", "5672")
		viper.SetDefault("RABBITMQ_USER", "guest")
		viper.SetDefault("RABBITMQ_PASS", "guest")
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
