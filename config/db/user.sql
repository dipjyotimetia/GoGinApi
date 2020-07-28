DROP TABLE users;

CREATE TABLE users
(
    uid SERIAL NOT NULL,
    name VARCHAR(100) NOT NULL,
    location VARCHAR(500) NOT NULL,
    age INT,
    CONSTRAINT userinfo_pkey PRIMARY KEY (uid)
);