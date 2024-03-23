-- name: InsertChampionship :one
INSERT INTO championship (rounds, family) Values ($1, $2) RETURNING *;

-- name: GetByID :one
SELECT * from championship where ID = $1;

-- name: UpdateChampionship :one
UPDATE championship set winnerID = $2, rounds = $3 where id = $1 RETURNING *;
