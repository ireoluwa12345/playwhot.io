-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX idx_users_username ON users(username);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_users_username;
-- +goose StatementEnd
