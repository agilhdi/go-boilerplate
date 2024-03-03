package init

import (
	"os"
	"strings"
	"time"

	"lolipad/boilerplate/pkg/util"

	"github.com/spf13/viper"
	log "go.uber.org/zap"
)

// Config for
type Config struct {
	API struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"api"`
	GRPC struct {
		Address string `mapstructure:"address"`
		Port    string `mapstructure:"port"`
	} `mapstructure:"grpc"`
	Context struct {
		Timeout int `mapstructure:"timeout"`
	} `mapstructure:"context"`
	Database struct {
		Pg struct {
			Host                  string        `mapstructure:"host"`
			Port                  int           `mapstructure:"port"`
			DBName                string        `mapstructure:"dbname"`
			User                  string        `mapstructure:"user"`
			Password              string        `mapstructure:"password"`
			SSLMode               string        `mapstructure:"sslmode"`
			MaxOpenConnection     int           `mapstructure:"max_open_connection"`
			MaxIdleConnection     int           `mapstructure:"max_idle_connection"`
			MaxConnectionLifetime time.Duration `mapstructure:"max_connection_lifetime"`
		} `mapstructure:"pg"`
		Redis struct {
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			Password string `mapstructure:"password"`
			DB       string `mapstructure:"db"`
		} `mapstructure:"redis"`
	} `mapstructure:"database"`
	Nsq struct {
		Host         string `mapstructure:"host"`
		Port         string `mapstructure:"port"`
		ProducerPort string `mapstructure:"producer_port"`
	} `mapstructure:"nsq"`
	Producer struct {
		EmailVerification   string `mapstructure:"email_verification"`
		EmailOtp            string `mapstructure:"email_otp"`
		EmailForgotPassword string `mapstructure:"email_forgot_password"`
	} `mapstructure:"producer"`
}

// setupMainConfig loads app config to viper
func setupMainConfig() (config *Config) {
	log.S().Info("Executing init/config")

	conf := false

	if util.IsDevelopmentEnv() {
		conf = true
		viper.AddRemoteProvider("consul", os.Getenv("CONFIG_ADDRESS"), os.Getenv("CONFIG_PATH"))
		viper.SetConfigType("json")
		err := viper.ReadRemoteConfig()
		if err != nil {
			log.S().Info("err: ", err)
		}
	}

	if util.IsProductionEnv() {
		conf = true
		log.S().Info("prod config")
		viper.SetConfigFile("config/app/production.json")
		err := viper.ReadInConfig()
		if err != nil {
			log.S().Info("err: ", err)
		}
	}

	if util.IsFileorDirExist("main.json") {
		conf = true
		log.S().Info("Local main.json file is found, now assigning it with default config")
		viper.SetConfigFile("main.json")
		err := viper.ReadInConfig()
		if err != nil {
			log.S().Info("err: ", err)
		}
	}

	if !conf {
		log.S().Fatal("Config is required")
	}

	viper.SetEnvPrefix(`app`)
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()

	err := viper.Unmarshal(&config)
	if err != nil {
		log.S().Fatal("err: ", err)
	}

	log.S().Info("Config APP_ENV: ", util.GetEnv())

	if !util.IsFileorDirExist("main.json") && !util.IsProductionEnv() {
		// open a goroutine to watch remote changes forever
		go func() {
			for {
				time.Sleep(time.Second * 5)

				err := viper.WatchRemoteConfig()
				if err != nil {
					log.S().Errorf("unable to read remote config: %v", err)
					continue
				}

				// unmarshal new config into our runtime config struct. you can also use channel
				// to implement a signal to notify the system of the changes
				viper.Unmarshal(&config)
			}
		}()
	}

	return
}
