-- V1__init.sql

CREATE TABLE owner (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    document VARCHAR(255) NOT NULL
);

CREATE TABLE wallet (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    balance INTEGER NOT NULL,
    owner_id VARCHAR(255) REFERENCES owner(id)
);
