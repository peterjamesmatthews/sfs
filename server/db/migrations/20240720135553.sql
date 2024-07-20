-- Modify "node" table
ALTER TABLE "public"."node" ADD CONSTRAINT "parent_not_self" CHECK ((parent IS NULL) OR (parent <> id));
-- Create index "unique_owner_name_parent" to table: "node"
CREATE UNIQUE INDEX "unique_owner_name_parent" ON "public"."node" ("owner", "name", (COALESCE(parent, '00000000-0000-0000-0000-000000000000'::uuid)));
