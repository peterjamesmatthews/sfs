-- Modify "user" table
ALTER TABLE "public"."user" ADD CONSTRAINT "user_name_key" UNIQUE ("name");
