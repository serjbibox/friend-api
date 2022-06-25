-- -------------------------------------------------------------
-- Database: friend_service
-- -------------------------------------------------------------
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS friend_id_seq;

-- Table Definition
CREATE TABLE "public"."friend_user" (
    "id" int4 NOT NULL DEFAULT nextval('friend_id_seq'::regclass),
    "date_added" timestamp,
    "phone" varchar,
    "login" varchar,
    "password" varchar,
    "name" varchar,
    "birth" varchar,
    "tg" varchar,
    "email" varchar,
    "login_status" BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY ("id")
);
