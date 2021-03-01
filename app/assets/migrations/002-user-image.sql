-- +migrate Up
alter table users
    ADD COLUMN image_url text not null default '';
ALTER TABLE users
    ADD CONSTRAINT unique_username UNIQUE (username);

-- +migrate Down
alter table users
    drop column image_url;
alter table users
    drop constraint unique_username;