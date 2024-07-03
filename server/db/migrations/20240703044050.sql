-- Modify "user" table
ALTER TABLE "public"."user" DROP COLUMN "name", ADD COLUMN "email" character varying NOT NULL, ADD CONSTRAINT "user_email_key" UNIQUE ("email");
