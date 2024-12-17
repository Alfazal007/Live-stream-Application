-- name: CreateUser :one
insert into users
    (id, username, password) 
        values ($1, $2, $3) returning *;

-- name: GetUserByName :one
select * from users where username=$1;

-- name: GetUseFromId :one
select * from users where id=$1;

