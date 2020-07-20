package handlers

import (
	"github.com/sirupsen/logrus"
)

type requestManager struct {
	log *logrus.Logger
}

// CreateHandlers : Returns the struct that carries all the handler functions
func CreateHandlers(log *logrus.Logger) *requestManager {
	return &requestManager{log}
}
