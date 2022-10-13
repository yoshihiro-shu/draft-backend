-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

INSERT INTO tags (name, description) VALUES ('tag-1', 'it is test tag-1');
INSERT INTO tags (name, description) VALUES ('tag-2', 'it is test tag-2');
INSERT INTO tags (name, description) VALUES ('tag-3', 'it is test tag-3');

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DELETE FROM articles WHERE name = 'tag-1';
DELETE FROM articles WHERE name = 'tag-2';
DELETE FROM articles WHERE name = 'tag-3';
-- +goose StatementEnd
