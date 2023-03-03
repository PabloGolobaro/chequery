package app

import "context"

func (a *Application) Stop() error {
	return a.router.Shutdown(context.Background())
}
