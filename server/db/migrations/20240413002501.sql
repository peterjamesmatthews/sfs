-- Create enum type "access_type"
CREATE TYPE "public"."access_type" AS ENUM ('read', 'write');
-- Create "access" table
CREATE TABLE "public"."access" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "user" uuid NOT NULL, "type" "public"."access_type" NOT NULL, "node" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "access_node" FOREIGN KEY ("node") REFERENCES "public"."node" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "access_user" FOREIGN KEY ("user") REFERENCES "public"."user" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
