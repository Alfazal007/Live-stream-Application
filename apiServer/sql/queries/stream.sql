-- name: CreateStream :one
insert into stream
    (id, admin_id)
        values ($1, $2) returning *;

-- name: StartStream :one
update stream
set started = true
where id=$1 and admin_id=$2
returning *;

-- name: EndStream :one
update stream
set ended = true
where id=$1 and admin_id=$2
returning *;

-- name: GetStreamFromId :one
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
    stream.admin_id = users.id;

-- name: GetStreamFromIdForWS :one
SELECT * FROM stream where id=$1;

