package bootstrap

import (
	"CSV/internal"
	"CSV/pkg/db"
)

func Execute() {

	db.Init()

	internal.Migrate()

}
