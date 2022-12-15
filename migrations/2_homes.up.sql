CREATE TABLE IF NOT EXISTS homes(
    id SERIAL PRIMARY KEY,
    PRICE BIGINT NOT NULL,
    DESCRIPTION VARCHAR NOT NULL,
    ADDRESS VARCHAR NOT NULL,
    agent_id BIGINT REFERENCES agents(id)
    
);