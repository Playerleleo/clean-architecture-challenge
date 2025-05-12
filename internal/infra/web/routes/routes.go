package routes

import (
	"net/http"

	"github.com/devfullcycle/20-CleanArch/internal/infra/web/handlers"
)

func SetupRoutes(orderHandler *handlers.OrderHandler) {
	http.HandleFunc("/order", orderHandler.List)
}
