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

CREATE TABLE IF NOT EXISTS accounts
(
    clientId     int          NOT NULL,
    currencyCode VARCHAR(100) NOT NULL,
    statusCode   VARCHAR(100) NOT NULL,
    balance      float8       NOT NULL,
    CONSTRAINT clientId_pkey PRIMARY KEY (clientId)
);

CREATE TABLE IF NOT EXISTS expense
(
    eid           SERIAL       NOT NULL,
    clientId      int          NOT NULL,
    expenseType   VARCHAR(100) NOT NULL,
    expenseAmount float8       NOT NULL,
    expenseDate   varchar(400) NOT NULL,
    CONSTRAINT expenseInfo_pkey PRIMARY KEY (eid)
);