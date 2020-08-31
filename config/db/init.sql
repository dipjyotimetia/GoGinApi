DROP TABLE IF EXISTS people;
DROP TABLE IF EXISTS videos;
DROP TABLE IF EXISTS expense;
DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS people
(
    id         bigserial not null
        constraint people_pkey
            primary key,
    first_name varchar(32),
    last_name  varchar(32),
    age        integer,
    email      varchar(256)
);

CREATE TABLE IF NOT EXISTS videos
(
    id          bigserial not null
        constraint videos_pkey
            primary key,
    title       varchar(100),
    description varchar(200),
    url         varchar(256)
        constraint videos_url_key
            unique,
    person_id   bigint,
    created_at  timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at  timestamp with time zone default CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS expense
(
    eid           SERIAL       NOT NULL,
    username      VARCHAR(100) NOT NULL,
    expenseType   VARCHAR(100) NOT NULL,
    expenseAmount float8       NOT NULL,
    expenseDate   varchar(400) NOT NULL,
    CONSTRAINT expenseInfo_pkey PRIMARY KEY (eid)
);

-- CREATE TABLE IF NOT EXISTS users
-- (
--     uid      SERIAL       NOT NULL,
--     name     VARCHAR(100) NOT NULL,
--     location VARCHAR(500) NOT NULL,
--     age      INT,
--     CONSTRAINT userinfo_pkey PRIMARY KEY (uid)
-- );


CREATE TABLE IF NOT EXISTS users
(
    id         serial PRIMARY KEY,
    name       VARCHAR(100)        NOT NULL,
    password   VARCHAR(355)        NOT NULL,
    email      VARCHAR(355) UNIQUE NOT NULL,
    created_on TIMESTAMP           NOT NULL default current_timestamp,
    updated_at TIMESTAMP           NOT NULL default current_timestamp
)