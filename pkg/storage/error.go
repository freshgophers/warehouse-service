package storage

import (
	"errors"
)

var ErrorNotFound = errors.New("store: no rows in result set")
