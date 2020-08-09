DROP TABLE expense;

CREATE TABLE expense
(
    eid           SERIAL       NOT NULL,
    username      VARCHAR(100) NOT NULL,
    expenseType   VARCHAR(100) NOT NULL,
    expenseAmount float8       NOT NULL,
    expenseDate   varchar(100) NOT NULL,
    CONSTRAINT expenseInfo_pkey PRIMARY KEY (eid)
);