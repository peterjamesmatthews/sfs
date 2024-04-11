-- Create "node" table
CREATE TABLE "public"."node" (
 "id" uuid NOT NULL DEFAULT gen_random_uuid(),
 "name" character varying NOT NULL,
 "owner" uuid NOT NULL,
 "parent" uuid NULL,
 PRIMARY KEY ("id"),
 CONSTRAINT "node_owner" FOREIGN KEY ("owner") REFERENCES "public"."user" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
 CONSTRAINT "node_parent" FOREIGN KEY ("parent") REFERENCES "public"."node" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
