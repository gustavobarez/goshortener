package handler

import (
	"goshortener/config"
)

var (
	logger *config.Logger
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
}
