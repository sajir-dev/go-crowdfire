## Install and Setup
Go version => 1.16

1. Clone the repo
2. cd into the directory *cd go-crowdfire-win*
3. wait for packages to dowload automatically or enter *go get -u ./...*
4. *go run main.go*

*env-variables*
DB_HOST=localhost
DB_PORT=27017
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_NAME=crowdfire

-- CREATE TABLE "users" (
--   "id" SERIAL PRIMARY KEY,
--   "name" varchar,
--   "email" varchar,
--   "password" varchar,
-- );

-- CREATE TABLE "posts" (
--   "created_at" varchar,
--   "id" SERIAL PRIMARY KEY,
--   "created_by" int,
--   "content" text
-- );

-- CREATE TABLE "following" (
--   "userid" varchar PRIMARY KEY,
--   "following" varchar[]
-- );

-- ALTER TABLE "posts" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("id");


-- ALTER TABLE "users" DROP COLUMN "following"

-- select * from following
-- select * from "following"
-- select * from "users"
-- SELECT "following" FROM "following" WHERE "userid" = 5::varchar
-- SELECT "following" FROM "following" WHERE "userid" = 5
-- SELECT "password" FROM "users" WHERE id = 5
-- UPDATE "following" SET "following" = array_append("following", 2) WHERE "userid" = 5::varchar;
-- UPDATE "following" SET "following"=[10,11] WHERE "userid" = 1::varchar

-- ALTER TABLE "users" ADD UNIQUE (email)
