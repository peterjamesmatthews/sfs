-- Modify "user" table
ALTER TABLE "public"."user" ALTER COLUMN "salt" DROP NOT NULL, ALTER COLUMN "hash" DROP NOT NULL;
