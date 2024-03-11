package handlers

import (
    "github.com/go-chi/chi/v5"
    chimiddle "github.com/go-chi/chi/v5/middleware"
    "github.com/z4rathustr4/goapi/internal/middleware"
)

func Handler(r *chi.Mux) {
    r.Use(chimiddle.StripSlashes)
    r.Route("/account", func(r chi.Router) {
        // middleware for /account route
       router.Use(middleware.Authorization)
       r.Get("/coins", GetCoinBalance) 
    }
}

