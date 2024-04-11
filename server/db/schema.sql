CREATE TABLE "user" (
  "id" UUID NOT NULL DEFAULT gen_random_uuid(),
  "name" VARCHAR NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "node" (
  "id" UUID NOT NULL DEFAULT gen_random_uuid(),
  "name" VARCHAR NOT NULL,
  "owner" UUID NOT NULL,
  "parent" UUID,
  PRIMARY KEY ("id"),
  CONSTRAINT "node_owner" FOREIGN KEY ("owner") REFERENCES "user" ("id") ON DELETE CASCADE,
  CONSTRAINT "node_parent" FOREIGN KEY ("parent") REFERENCES "node" ("id") ON DELETE CASCADE
);
