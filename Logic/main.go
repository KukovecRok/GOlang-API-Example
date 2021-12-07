package Logic

import (
	"todorokvaja1/DB"
)

type Controller struct {
	db     DB.DB
	secret []byte
}

func NewController(db DB.DB, secret []byte) *Controller {
	return &Controller{
		db:     db,
		secret: secret,
	}
}
