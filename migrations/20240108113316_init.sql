-- +goose Up
-- +goose StatementBegin
CREATE TYPE statuses AS ENUM (
  'InProgress',
  'Done',
  'Error'
);

CREATE TABLE tasks (
  "id" BIGSERIAL PRIMARY KEY,
  "company_id" int NOT NULL,
  "status" statuses NOT NULL
);

CREATE TABLE tasks_result (
  "id" BIGSERIAL PRIMARY KEY,
  "task_id" bigint UNIQUE NOT NULL,
  "result" JSON NOT NULL
);
ALTER TABLE "tasks_result" ADD FOREIGN KEY ("task_id") REFERENCES "tasks" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks_result;
DROP TABLE tasks;
DROP TYPE statuses;
-- +goose StatementEnd
