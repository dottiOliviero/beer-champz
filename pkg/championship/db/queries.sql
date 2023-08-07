-- name: InsertChampionship :one
INSERT INTO championship (rounds) Values ($1) RETURNING *;

-- name: GetByID :one
SELECT * from championship where ID = $1;

-- name: UpdateChampionship :one
UPDATE championship set winnerID = $2, rounds = $3 where id = $1 RETURNING *;
