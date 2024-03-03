package init

import (
	"github.com/spf13/viper"
	log "go.uber.org/zap"
)

// setupAuthHelper inits
func setupAuthHelper() {
	if !viper.IsSet("auth.access_token_secret") || viper.GetString("auth.access_token_secret") == "" {
		log.S().Fatal("auth.access_token_secret can not be empty for better security on auth")
	}

	if !viper.IsSet("auth.refresh_token_secret") || viper.GetString("auth.refresh_token_secret") == "" {
		log.S().Fatal("auth.refresh_token_secret can not be empty for better security on auth")
	}
}
