CREATE DATABASE db;
\connect db;
CREATE SCHEMA test;
SET search_path to test;
CREATE TABLE db.test.entity
(
    id          SERIAL PRIMARY KEY,
    name        TEXT,
    description TEXT,
    properties  JSONB
);

INSERT INTO db.test.entity (id, name, description, properties) VALUES (1, 'test', 'any desc', '{"amounts": [{"amount": 999991, "image_url": "https://image.com/example", "description": "12345"}, {"amount": 20000, "image_url": "https://image.com/example", "description": "3456"}]}');