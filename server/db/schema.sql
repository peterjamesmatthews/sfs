CREATE TABLE "user" (
  "id" UUID NOT NULL UNIQUE DEFAULT gen_random_uuid(),
  PRIMARY KEY ("id"),
  "name" VARCHAR NOT NULL
);

CREATE TABLE "node" (
  "id" UUID NOT NULL UNIQUE DEFAULT gen_random_uuid(),
  PRIMARY KEY ("id"),
  "name" VARCHAR NOT NULL,
  "owner" UUID NOT NULL,
  CONSTRAINT "node_owner" FOREIGN KEY ("owner") REFERENCES "user" ("id") ON DELETE CASCADE,
  "parent" UUID,
  CONSTRAINT "node_parent" FOREIGN KEY ("parent") REFERENCES "node" ("id") ON DELETE CASCADE,
  -- no nodes may have the same name, owner, and parent
  CONSTRAINT "node_name_owner_parent" UNIQUE ("name", "owner", "parent")
);

CREATE TABLE "file" (
  "id" UUID NOT NULL UNIQUE DEFAULT gen_random_uuid(),
  PRIMARY KEY ("id"),
  "node" UUID NOT NULL,
  CONSTRAINT "file_node" FOREIGN KEY ("node") REFERENCES "node" ("id") ON DELETE CASCADE,
  "content" TEXT NOT NULL
);

CREATE TYPE "access_type" AS ENUM ('read', 'write');

CREATE TABLE "access" (
  "id" UUID NOT NULL UNIQUE DEFAULT gen_random_uuid(),
  PRIMARY KEY ("id"),
  "user" UUID NOT NULL,
  CONSTRAINT "access_user" FOREIGN KEY ("user") REFERENCES "user" ("id") ON DELETE CASCADE,
  "type" "access_type" NOT NULL,
  "node" UUID NOT NULL,
  CONSTRAINT "access_node" FOREIGN KEY ("node") REFERENCES "node" ("id") ON DELETE CASCADE,
  -- no users may have duplicate access to the same node
  CONSTRAINT "access_user_type_node" UNIQUE ("user", "type", "node")
);
