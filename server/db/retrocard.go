package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rob-lowcock/pioneer2023/models"
)

type Retrocard struct {
	Db *pgxpool.Pool
}

func (r *Retrocard) GetActiveCards() ([]models.Retrocard, error) {
	rows, err := r.Db.Query(context.Background(), `SELECT id, title, col, active FROM retrocards WHERE active = true`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	retrocards := []models.Retrocard{}
	for rows.Next() {
		retrocard := models.Retrocard{}
		err := rows.Scan(&retrocard.ID, &retrocard.Title, &retrocard.Column, &retrocard.Active)
		if err != nil {
			return nil, err
		}
		retrocards = append(retrocards, retrocard)
	}

	return retrocards, nil
}

func (r *Retrocard) CreateRetrocard(retrocard models.Retrocard) (string, error) {
	id := ""
	retrocard.Tidy()
	err := retrocard.Validate()

	if err != nil {
		return id, err
	}

	err = r.Db.QueryRow(context.Background(), `INSERT INTO retrocards (title, col, active) VALUES ($1, $2, $3) RETURNING id`, retrocard.Title, retrocard.Column, retrocard.Active).Scan(&id)

	return id, err
}
