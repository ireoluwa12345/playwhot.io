-- +goose Up
-- +goose StatementBegin
ALTER TABLE rooms ADD COLUMN name VARCHAR(255) NOT NULL DEFAULT '';
ALTER TABLE rooms ADD COLUMN max_players INT NOT NULL DEFAULT 4;
ALTER TABLE rooms ADD COLUMN is_private BOOLEAN NOT NULL DEFAULT FALSE;
ALTER TABLE rooms ADD COLUMN password VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE rooms DROP COLUMN IF EXISTS password;
ALTER TABLE rooms DROP COLUMN IF EXISTS is_private;
ALTER TABLE rooms DROP COLUMN IF EXISTS max_players;
ALTER TABLE rooms DROP COLUMN IF EXISTS name;
-- +goose StatementEnd