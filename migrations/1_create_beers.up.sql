CREATE TABLE IF NOT EXISTS beers (
	id serial PRIMARY KEY,
	name       varchar(50),
	style      varchar(50),
	sub_style  varchar(50),
	abv        varchar(10),
	short_desc varchar(50),
	brewery    varchar(50),
	image      varchar(100),
	score      int
)
