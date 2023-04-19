CREATE TABLE IF NOT EXISTS fortunes (
    id SERIAL PRIMARY KEY,
    result TEXT NOT NULL
);

INSERT INTO fortunes (result) VALUES
    ('大吉'),
    ('中吉'),
    ('小吉'),
    ('吉'),
    ('末吉')
;
