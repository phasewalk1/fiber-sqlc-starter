package handlers

import (
	"github.com/jackc/pgx/v5"
	"github.com/phasewalk1/fiber-sqlc-starter/database"
)

type Handler struct {
	conn      *pgx.Conn
	queries   *database.Queries
	validator *Validator
}
