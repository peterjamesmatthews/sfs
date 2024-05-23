CREATE TABLE "user" (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR NOT NULL UNIQUE
);

CREATE TABLE "node" (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  owner UUID NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
  name VARCHAR NOT NULL,
  parent UUID REFERENCES "node" (id) ON DELETE CASCADE
);

CREATE TABLE "file" (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  node UUID NOT NULL REFERENCES node (id) ON DELETE CASCADE,
  content BYTEA NOT NULL DEFAULT ''
);

CREATE TYPE "access_type" AS ENUM (
  'read',
  'write'
);

CREATE TABLE "access" (
  accessor UUID NOT NULL REFERENCES "user" (id) ON DELETE CASCADE,
  "target" UUID NOT NULL REFERENCES "node" (id) ON DELETE CASCADE,
  type "access_type" NOT NULL,
  PRIMARY KEY (accessor, "target")
);
