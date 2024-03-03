package init

import (
	"database/sql"
	"os"

	"lolipad/boilerplate/internal/db"
	"lolipad/boilerplate/pkg/util"

	"github.com/spf13/viper"
	log "go.uber.org/zap"
)

// ConnectToPGServer is a function to init PostgreSQL connection
func ConnectToPGServer() (*sql.DB, error) {
	if util.IsProductionEnv() && (!viper.IsSet("database.pg.password") || viper.GetString("database.pg.password") == "") {
		log.S().Fatal("database.pg.password can not be empty!")
	}

	dbpg, err := db.CreatePGConnection(map[string]string{
		"host":     viper.GetString(`database.pg.host`),
		"port":     viper.GetString(`database.pg.port`),
		"user":     viper.GetString(`database.pg.user`),
		"password": viper.GetString(`database.pg.password`),
		"dbname":   viper.GetString(`database.pg.dbname`),
		"sslmode":  viper.GetString(`database.pg.sslmode`),
	})

	if err != nil {
		os.Exit(1)
	}

	return dbpg, err
}
