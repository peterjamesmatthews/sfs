-- Create "file" table
CREATE TABLE "public"."file" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "node" uuid NOT NULL, "content" text NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "file_node" FOREIGN KEY ("node") REFERENCES "public"."node" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
