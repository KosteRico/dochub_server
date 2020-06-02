package main

import "github.com/rs/cors"

func corsConfig(allowedUrl ...string) *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   allowedUrl,
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Filename"},
		AllowCredentials: true,
		Debug:            true,
	})
}
