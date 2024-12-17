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
        values ($1, $2) returning id, admin_id, started, ended
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
		&i.Started,
		&i.Ended,
	)
	return i, err
}

const endStream = `-- name: EndStream :one
update stream
set ended = true
where id=$1
returning id, admin_id, started, ended
`

func (q *Queries) EndStream(ctx context.Context, id string) (Stream, error) {
	row := q.db.QueryRowContext(ctx, endStream, id)
	var i Stream
	err := row.Scan(
		&i.ID,
		&i.AdminID,
		&i.Started,
		&i.Ended,
	)
	return i, err
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

const startStream = `-- name: StartStream :one
update stream
set started = true
where id=$1
returning id, admin_id, started, ended
`

func (q *Queries) StartStream(ctx context.Context, id string) (Stream, error) {
	row := q.db.QueryRowContext(ctx, startStream, id)
	var i Stream
	err := row.Scan(
		&i.ID,
		&i.AdminID,
		&i.Started,
		&i.Ended,
	)
	return i, err
}
