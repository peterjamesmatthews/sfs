-- Modify "file" table
ALTER TABLE "public"."file" ADD CONSTRAINT "file_node_key" UNIQUE ("node");
