// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package dbquery

import (
	"context"

	"github.com/google/uuid"
)

const createCampaign = `-- name: CreateCampaign :one
INSERT INTO campaigns (
    name, description
) VALUES (
    $1, $2
)
ON CONFLICT(name) DO UPDATE
    SET name = excluded.name
RETURNING id, name, description, adserver_id, created_at, updated_at
`

type CreateCampaignParams struct {
	Name        string
	Description string
}

func (q *Queries) CreateCampaign(ctx context.Context, arg CreateCampaignParams) (Campaign, error) {
	row := q.db.QueryRow(ctx, createCampaign, arg.Name, arg.Description)
	var i Campaign
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.AdserverID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCampaign = `-- name: GetCampaign :one
SELECT id, name, description, adserver_id, created_at, updated_at
FROM campaigns
WHERE id = $1
`

func (q *Queries) GetCampaign(ctx context.Context, id *uuid.UUID) (Campaign, error) {
	row := q.db.QueryRow(ctx, getCampaign, id)
	var i Campaign
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.AdserverID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCampaign = `-- name: UpdateCampaign :one
UPDATE campaigns
SET adserver_id = $1
WHERE id = $2
RETURNING id, name, description, adserver_id, created_at, updated_at
`

type UpdateCampaignParams struct {
	AdserverID *uuid.UUID
	ID         *uuid.UUID
}

func (q *Queries) UpdateCampaign(ctx context.Context, arg UpdateCampaignParams) (Campaign, error) {
	row := q.db.QueryRow(ctx, updateCampaign, arg.AdserverID, arg.ID)
	var i Campaign
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.AdserverID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
