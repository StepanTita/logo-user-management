-- +migrate Up
create table users
(
    id       bigserial,
    username text not null unique,
    name     text not null default '',
    surname  text not null default '',
    salt     text not null default '',
    email    text not null,
    password text not null,
    PRIMARY KEY (id)
);

-- +migrate Down
drop table users cascade;