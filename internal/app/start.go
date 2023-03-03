package app

import "github.com/pablogolobaro/chequery/internal/config"

func (a *Application) Start(conf config.Config) error {
	return a.router.Start(conf.HttpHost + ":" + conf.HttpPort)
}
