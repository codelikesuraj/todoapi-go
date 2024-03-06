CREATE DATABASE IF NOT EXISTS todoapigo;

USE todoapigo;

DROP TABLE IF EXISTS todos;

CREATE TABLE todos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    completed BOOLEAN NOT NULL DEFAULT false,
    due TIMESTAMP NOT NULL,
    todo VARCHAR(64) NOT NULL
);