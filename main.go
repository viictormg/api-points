package main

import (
	"api-points/config"
	"api-points/features"
	"api-points/repository"
	"api-points/server"

	"log"
)

func run() error {
	cfg, err := config.NewConfig()

	if err != nil {
		return err
	}

	db, err := config.InitDB(*cfg)

	if err != nil {
		return err
	}

	Repository := repository.NewRepository(db.Database("leal_points"))
	Feature := features.NewFeature(Repository)
	srv := server.NewServer(cfg, Feature)

	return srv.Run()

}
func main() {

	err := run()

	if err != nil {
		log.Panic(err)
	}

}
