-- Create "user" table
CREATE TABLE "public"."user" (
 "id" uuid NOT NULL DEFAULT gen_random_uuid(),
 "name" character varying NOT NULL,
 PRIMARY KEY ("id")
);
