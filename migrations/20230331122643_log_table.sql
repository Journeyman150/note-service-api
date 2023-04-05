-- +goose Up
create table log (
      id          bigserial primary key,
      note_id     bigint references note(id),
      msg         text,
      created_at  timestamp not null default now()
);

-- +goose Down
drop table log;