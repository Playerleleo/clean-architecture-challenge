package routes

import (
	"net/http"

	"github.com/leonardogomesdossantos/clean-architecture-challenge/internal/infra/web/handlers"
)

func SetupRoutes(orderHandler *handlers.OrderHandler) {
	http.HandleFunc("/order", orderHandler.List)
}
