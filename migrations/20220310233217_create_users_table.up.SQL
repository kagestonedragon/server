CREATE SCHEMA IF NOT EXISTS "user";

CREATE TABLE "user"."users" (
    "id" serial8,
    "name" varchar(255) NOT NULL,
    "is_active" bool NOT NULL DEFAULT true,
    PRIMARY KEY ("id")
);
