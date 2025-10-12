-- V1__init.sql

CREATE TABLE owners (
    id VARCHAR(255) PRIMARY KEY,
    owner_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    document VARCHAR(255) NOT NULL
);

CREATE TABLE wallets (
    id VARCHAR(255) PRIMARY KEY,
    wallet_name VARCHAR(255) NOT NULL,
    balance INTEGER NOT NULL,
    owner_id VARCHAR(255) REFERENCES owners(id)
);
