package bootstrap

import (
	"auctionsystem/pkg/logger"
	"fmt"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`

	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`

	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPass     string `mapstructure:"REDIS_PASS"`
	RedisDatabase int    `mapstructure:"REDIS_DATABASE"`

	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func (env *Env) String() string {
	return fmt.Sprintf("Env{AppEnv=%s, ServerAddress=%s, ContextTimeout=%d, DBHost=%s, DBPort=%s, DBUser=%s, DBPass=%s, DBName=%s, AccessTokenExpiryHour=%d, RefreshTokenExpiryHour=%d, AccessTokenSecret=%s, RefreshTokenSecret=%s}",
		env.AppEnv, env.ServerAddress, env.ContextTimeout, env.DBHost, env.DBPort, env.DBUser, env.DBPass, env.DBName, env.AccessTokenExpiryHour, env.RefreshTokenExpiryHour, env.AccessTokenSecret, env.RefreshTokenSecret)
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Logger.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		logger.Logger.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		logger.Logger.Println("The App is running in development env")
	}

	return &env
}
