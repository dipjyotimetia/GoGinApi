DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS expense;

CREATE TABLE IF NOT EXISTS users
(
    id         serial PRIMARY KEY,
    name       VARCHAR(100)        NOT NULL,
    password   VARCHAR(355)        NOT NULL,
    email      VARCHAR(355) UNIQUE NOT NULL,
    created_on TIMESTAMP           NOT NULL default current_timestamp,
    updated_at TIMESTAMP           NOT NULL default current_timestamp
);

INSERT INTO users
VALUES (1, 'test1', '$2a$10$tqcfZ6IKzeXnFH8VoauRouWcfKRHDTqRyWg3BKj/xUKiM.L.xfNAC', 'test1@gmail.com',
        current_timestamp, current_timestamp);

CREATE TABLE IF NOT EXISTS accounts
(
    accountId    SERIAL PRIMARY KEY,
    currencyCode VARCHAR(100) NOT NULL,
    statusCode   VARCHAR(100) NOT NULL,
    balance      float8       NOT NULL,
    clientId     int REFERENCES users (id)
);

INSERT INTO accounts
VALUES (1, 'AUD', 'ACTIVE', 200.50, 1);

CREATE TABLE IF NOT EXISTS expense
(
    eid           SERIAL PRIMARY KEY,
    expenseType   VARCHAR(100) NOT NULL,
    expenseAmount float8       NOT NULL,
    expenseDate   varchar(400) NOT NULL,
    clientId      int REFERENCES users (id)
);

INSERT INTO expense
VALUES (1, 'Test1', 20.50, '12/12/2019', 1);