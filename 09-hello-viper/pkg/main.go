package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	configuration "potato/hello-viper/configuration"
)

func init() {
	profile := os.Getenv("GO_ACTIVE_PROFILE")
	if len(profile) <= 0 {
		profile = "local"
	}

	log.Println("ACTIVE PROFILE: ", profile)

	// init viper
	viper.AddConfigPath(".")
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Could not load configuration profile(%v)", profile)
		panic(err)
	}

	if err := viper.Unmarshal(&configuration.RuntimeConf); err != nil {
		log.Fatalln(err)
		panic(err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config file changed:", e.Name)
		if err := viper.ReadInConfig(); err != nil {
			log.Println(err)
			return
		}

		if err := viper.Unmarshal(&configuration.RuntimeConf); err != nil {
			log.Println(err)
			return
		}
	})

	viper.WatchConfig()
}

func main() {
	log.Println(configuration.RuntimeConf)
}
