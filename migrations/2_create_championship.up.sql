CREATE TABLE IF NOT EXISTS championship (
    id serial PRIMARY KEY,
    winnerID integer,
    rounds jsonb
)
