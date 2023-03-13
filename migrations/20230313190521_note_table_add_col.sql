-- +goose Up
alter table note
    add column email text;

-- +goose Down
alter table note
    drop column email;


