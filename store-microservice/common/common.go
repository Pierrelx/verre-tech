package common

import (
	"encoding/json"
	"os"

	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

//Configuration comporte les éléments de configuration
type Configuration struct {
	LogFilename   string `json:"logFilename"`
	LogMaxSize    int    `json:"logMaxSize"`
	LogMaxBackups int    `json:"logMaxBackups"`
	LogMaxAge     int    `json:"logMaxAge"`

	DbAddress  string `json:"dbAddress"`
	DbName     string `json:"dbName"`
	DbUserName string `json:"dbUserName"`
	DbPassword string `json:"dbPassword"`
	DbPort     string `json:"dbPort"`
}

// Configuration est partagé
var (
	Config *Configuration
)

// LoadConfig permet de charger le fichier de configuration
func LoadConfig() error {
	file, err := os.Open("config/config.json")
	if err != nil {
		return err
	}

	Config = new(Configuration)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	log.SetOutput(&lumberjack.Logger{
		Filename:   Config.LogFilename,
		MaxSize:    Config.LogMaxSize,    //MB
		MaxBackups: Config.LogMaxBackups, //Nombre de backups
		MaxAge:     Config.LogMaxAge,     //En jours
	})
	log.SetLevel(log.DebugLevel)

	log.SetFormatter(&log.JSONFormatter{})

	return nil
}
