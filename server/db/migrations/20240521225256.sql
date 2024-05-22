-- Create enum type "access_type"
CREATE TYPE "public"."access_type" AS ENUM ('read', 'write');
-- Create "user" table
CREATE TABLE "public"."user" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "name" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "node" table
CREATE TABLE "public"."node" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "owner" uuid NOT NULL, "name" character varying NOT NULL, "parent" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "node_owner_fkey" FOREIGN KEY ("owner") REFERENCES "public"."user" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "node_parent_fkey" FOREIGN KEY ("parent") REFERENCES "public"."node" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "access" table
CREATE TABLE "public"."access" ("accessor" uuid NOT NULL, "target" uuid NOT NULL, "type" "public"."access_type" NOT NULL, PRIMARY KEY ("accessor", "target"), CONSTRAINT "access_accessor_fkey" FOREIGN KEY ("accessor") REFERENCES "public"."user" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "access_target_fkey" FOREIGN KEY ("target") REFERENCES "public"."node" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "file" table
CREATE TABLE "public"."file" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "node" uuid NOT NULL, "content" bytea NOT NULL DEFAULT '\x', PRIMARY KEY ("id"), CONSTRAINT "file_node_fkey" FOREIGN KEY ("node") REFERENCES "public"."node" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
