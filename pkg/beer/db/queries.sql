-- name: InsertBeer :one
INSERT INTO beers (name, style, sub_style, abv, short_desc, brewery, image, score) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;


-- name: GetAll :many
SELECT * FROM beers;
