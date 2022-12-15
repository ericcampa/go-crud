CREATE TABLE IF NOT EXISTS agents(
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL
);

INSERT INTO agents values (1, 'agente_1');