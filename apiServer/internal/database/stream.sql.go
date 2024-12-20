// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: stream.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createStream = `-- name: CreateStream :one
insert into stream
    (id, admin_id)
        values ($1, $2) returning id, admin_id, created_at, started, ended
`

type CreateStreamParams struct {
	ID      string
	AdminID uuid.NullUUID
}

func (q *Queries) CreateStream(ctx context.Context, arg CreateStreamParams) (Stream, error) {
	row := q.db.QueryRowContext(ctx, createStream, arg.ID, arg.AdminID)
	var i Stream
	err := row.Scan(
		&i.ID,
		&i.AdminID,
		&i.CreatedAt,
		&i.Started,
		&i.Ended,
	)
	return i, err
}

const endStream = `-- name: EndStream :one
update stream
set ended = true
where id=$1 and admin_id=$2
returning id, admin_id, created_at, started, ended
`

type EndStreamParams struct {
	ID      string
	AdminID uuid.NullUUID
}

func (q *Queries) EndStream(ctx context.Context, arg EndStreamParams) (Stream, error) {
	row := q.db.QueryRowContext(ctx, endStream, arg.ID, arg.AdminID)
	var i Stream
	err := row.Scan(
		&i.ID,
		&i.AdminID,
		&i.CreatedAt,
		&i.Started,
		&i.Ended,
	)
	return i, err
}

const get10LatestStream = `-- name: Get10LatestStream :many
SELECT s.id, u.username
FROM stream s, users u
where s.started=true and s.ended=false and s.admin_id == u.id
ORDER BY s.created_at DESC
limit 10
`

type Get10LatestStreamRow struct {
	ID       string
	Username string
}

func (q *Queries) Get10LatestStream(ctx context.Context) ([]Get10LatestStreamRow, error) {
	rows, err := q.db.QueryContext(ctx, get10LatestStream)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Get10LatestStreamRow
	for rows.Next() {
		var i Get10LatestStreamRow
		if err := rows.Scan(&i.ID, &i.Username); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getStreamFromId = `-- name: GetStreamFromId :one
SELECT
    stream.id AS stream_id,
    stream.admin_id,
    stream.started,
    stream.ended,
    users.username AS admin_name
FROM
    stream
JOIN
    users
ON
    stream.admin_id = users.id
`

type GetStreamFromIdRow struct {
	StreamID  string
	AdminID   uuid.NullUUID
	Started   sql.NullBool
	Ended     sql.NullBool
	AdminName string
}

func (q *Queries) GetStreamFromId(ctx context.Context) (GetStreamFromIdRow, error) {
	row := q.db.QueryRowContext(ctx, getStreamFromId)
	var i GetStreamFromIdRow
	err := row.Scan(
		&i.StreamID,
		&i.AdminID,
		&i.Started,
		&i.Ended,
		&i.AdminName,
	)
	return i, err
}

const getStreamFromIdForWS = `-- name: GetStreamFromIdForWS :one
SELECT id, admin_id, created_at, started, ended FROM stream where id=$1
`

func (q *Queries) GetStreamFromIdForWS(ctx context.Context, id string) (Stream, error) {
	row := q.db.QueryRowContext(ctx, getStreamFromIdForWS, id)
	var i Stream
	err := row.Scan(
		&i.ID,
		&i.AdminID,
		&i.CreatedAt,
		&i.Started,
		&i.Ended,
	)
	return i, err
}

const startStream = `-- name: StartStream :one
update stream
set started = true
where id=$1 and admin_id=$2
returning id, admin_id, created_at, started, ended
`

type StartStreamParams struct {
	ID      string
	AdminID uuid.NullUUID
}

func (q *Queries) StartStream(ctx context.Context, arg StartStreamParams) (Stream, error) {
	row := q.db.QueryRowContext(ctx, startStream, arg.ID, arg.AdminID)
	var i Stream
	err := row.Scan(
		&i.ID,
		&i.AdminID,
		&i.CreatedAt,
		&i.Started,
		&i.Ended,
	)
	return i, err
}
