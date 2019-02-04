package main

import (
	"github.com/FelipeUmpierre/investiments-consumer/pkg/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {
	conf, err := config.New()
	failOnErr(err, "failed to load the environment variables")

	initDB(conf)
}

func initDB(conf *config.Config) *sqlx.DB {
	eventstore, err := sqlx.Open("postgres", conf.API)
	failOnErr(err, "failed to open connection with eventstore database")

	return eventstore
}

func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatal().Err(err).Msg(msg)
	}
}
