package handler

import (
	"database/sql"

	"social-network/internal/repository"
	"social-network/internal/service"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(db *sql.DB) *Handler {
	data := repository.Repository{
		Db: db,
	}

	service := service.Service{
		Database: &data,
	}

	return &Handler{
		Service: &service,
	}
}