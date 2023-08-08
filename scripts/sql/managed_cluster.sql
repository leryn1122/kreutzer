DROP TABLE IF EXISTS "managed_cluster";
CREATE TABLE "managed_cluster"
(
  "id"            SERIAL4      NOT NULL,
  "creation_time" TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "modified_time" TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "name"          VARCHAR(30)  NOT NULL,
  "url"           VARCHAR(256) NOT NULL,
  "enabled"       BOOL         NOT NULL DEFAULT TRUE,
  "cert"          TEXT         NOT NULL,
  "username"      VARCHAR(30)  NOT NULL,
  "token"         TEXT         NOT NULL,

  PRIMARY KEY ("id"),
  UNIQUE ("name")
);

COMMENT ON TABLE "managed_cluster" IS 'Managed kubernetes cluster';

COMMENT ON COLUMN "managed_cluster"."id" IS 'PK';
COMMENT ON COLUMN "managed_cluster"."creation_time" IS 'Creation timestamp';
COMMENT ON COLUMN "managed_cluster"."modified_time" IS 'Modified timestamp';
COMMENT ON COLUMN "managed_cluster"."name" IS 'Cluster name';
COMMENT ON COLUMN "managed_cluster"."url" IS 'Kubernetes API Server endpoint';
COMMENT ON COLUMN "managed_cluster"."enabled" IS 'Enabled or not';
COMMENT ON COLUMN "managed_cluster"."cert" IS 'Kubernetes certificate';
COMMENT ON COLUMN "managed_cluster"."username" IS 'Username access to kubernetes cluster';
COMMENT ON COLUMN "managed_cluster"."token" IS 'Token for access to kubernetes cluster';
