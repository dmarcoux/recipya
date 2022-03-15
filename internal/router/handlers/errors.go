package handlers

import (
	"net/http"

	"github.com/reaper47/recipya/internal/logger"
	"github.com/reaper47/recipya/internal/templates"
)

func showErrorPage(w http.ResponseWriter, message string, err error) {
	logger.Sanitize(message, err.Error())
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusInternalServerError)
	templates.Render(w, "error-500", message)
}
