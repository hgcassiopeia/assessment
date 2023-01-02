CREATE TABLE IF NOT EXISTS expenses (
    id SERIAL PRIMARY KEY,
    title TEXT,
    amount FLOAT,
    note TEXT,
    tags TEXT[]
);

INSERT INTO expenses (title, amount, note, tags) VALUES ('Isakaya Bangna', 899, 'central bangna', ARRAY['food','beverage']);