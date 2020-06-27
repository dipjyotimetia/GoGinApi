drop table people;
drop table videos;

create table people
(
    id         bigserial not null
        constraint people_pkey
            primary key,
    first_name varchar(32),
    last_name  varchar(32),
    age        integer,
    email      varchar(256)
);

create table videos
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

