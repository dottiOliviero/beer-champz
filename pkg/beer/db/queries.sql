-- name: InsertBeer :one
INSERT INTO beers (name, style, sub_style, abv, short_desc, brewery, image, score) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;


-- name: GetAll :many
SELECT * FROM beers order by score DESC;


-- name: UpdateBeerScore :one
UPDATE beers SET score = score + 1 where id = $1 RETURNING *;
