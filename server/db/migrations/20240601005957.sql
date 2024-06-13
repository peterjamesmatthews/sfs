-- Create "access_token" table
CREATE TABLE "public"."access_token" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "owner" uuid NOT NULL, "hash" bytea NOT NULL, "creation" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, "expiration" timestamp NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "access_token_owner_fkey" FOREIGN KEY ("owner") REFERENCES "public"."user" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "refresh_token" table
CREATE TABLE "public"."refresh_token" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "owner" uuid NOT NULL, "hash" bytea NOT NULL, "creation" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, "expiration" timestamp NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "refresh_token_owner_fkey" FOREIGN KEY ("owner") REFERENCES "public"."user" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);