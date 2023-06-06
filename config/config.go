package config

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	Configuration struct {
		Host        string `mapstructure:"host"`
		Port        string `mapstructure:"port"`
		Username    string `mapstructure:"username"`
		Password    string `mapstructure:"password"`
		DBName      string `mapstructure:"dbName"`
		ServicePort string `mapstructure:"servicePort"`
	}
	Conn struct {
		PostgreCon *gorm.DB
	}
)

var (
	connection *Conn
	Conf       *Configuration
)

func getEnvConfig() *Configuration {

	envFlag := flag.String("env", "dev", "state the active profile, default staging")
	flag.Parse()

	env := *envFlag

	viper.AddConfigPath("./config")
	viper.SetConfigName(fmt.Sprintf("config.%v", env))
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Cannot load configuration file: %v", err)
	}

	conf := &Configuration{}
	err = viper.Unmarshal(conf)
	if err != nil {
		log.Fatalf("Cannot unmarshall config: %v", err)
	}
	return conf
}

func GetConnection() *Conn {
	if connection != nil {
		return connection
	}
	Conf = getEnvConfig()
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v TimeZone=Asia/Jakarta sslmode=disable",
		Conf.Host, Conf.Username, Conf.Password, Conf.DBName, Conf.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error opening postgres connection: %v", err)
	}

	fmt.Println("Connected to database!")

	connection = &Conn{
		PostgreCon: db,
	}
	return connection
}
