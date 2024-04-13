-- Create index "access_user_type_node" to table: "access"
CREATE UNIQUE INDEX "access_user_type_node" ON "public"."access" ("user", "type", "node");
-- Create index "node_name_owner_parent" to table: "node"
CREATE UNIQUE INDEX "node_name_owner_parent" ON "public"."node" ("name", "owner", "parent");
