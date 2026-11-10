CREATE SCHEMA "todo";

CREATE TYPE "todo"."todo_status" AS ENUM ('TODO_STATUS_ACTIVE', 'TODO_STATUS_DELETED');

CREATE TABLE "todo"."todo" (
    "id" uuid NOT NULL,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    "deleted_at" timestamptz NULL,
    "title" varchar NOT NULL DEFAULT '',
    "content" varchar NOT NULL DEFAULT '',
    "completed" boolean NOT NULL DEFAULT false,
    "status" "todo"."todo_status" NOT NULL DEFAULT 'TODO_STATUS_ACTIVE',
    PRIMARY KEY ("id")
);

CREATE INDEX "todo_updated_at" ON "todo"."todo" ("updated_at");
CREATE INDEX "todo_deleted_at" ON "todo"."todo" ("deleted_at");
CREATE INDEX "todo_status" ON "todo"."todo" ("status");
