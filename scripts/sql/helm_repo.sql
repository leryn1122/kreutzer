DROP TABLE IF EXISTS "helm_repo";
CREATE TABLE "helm_repo"
(
  "id"                       SERIAL4      NOT NULL,
  "creation_time"            TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "modified_time"            TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "name"                     VARCHAR(64)  NOT NULL,
  "url"                      VARCHAR(64)  NOT NULL,
  "last_sync_time"           TIMESTAMP    NULL,
  "username"                 VARCHAR(64)  NULL,
  "password"                 VARCHAR(64)  NULL,
  "pass_credentials_all"     BOOL                  DEFAULT FALSE,
  "insecure_skip_tls_verify" BOOL                  DEFAULT FALSE,
  "cert_file"                VARCHAR(128) NULL,
  "key_file"                 VARCHAR(128) NULL,
  "ca_file"                  VARCHAR(128) NULL,
  PRIMARY KEY ("id")
);

COMMENT ON TABLE "helm_repo" IS 'Helm Repo';

COMMENT ON COLUMN "helm_repo"."id" IS 'PK';
COMMENT ON COLUMN "helm_repo"."creation_time" IS 'Creation timestamp';
COMMENT ON COLUMN "helm_repo"."modified_time" IS 'Modified timestamp';
COMMENT ON COLUMN "helm_repo"."name" IS 'Repo name';
COMMENT ON COLUMN "helm_repo"."url" IS 'Repo URL';
COMMENT ON COLUMN "helm_repo"."last_sync_time" IS 'Last update time';
COMMENT ON COLUMN "helm_repo"."username" IS 'Username';
COMMENT ON COLUMN "helm_repo"."password" IS 'Password';
COMMENT ON COLUMN "helm_repo"."pass_credentials_all" IS '';
COMMENT ON COLUMN "helm_repo"."insecure_skip_tls_verify" IS '';
COMMENT ON COLUMN "helm_repo"."cert_file" IS 'Certificate file';
COMMENT ON COLUMN "helm_repo"."key_file" IS 'Certificate key file';
COMMENT ON COLUMN "helm_repo"."ca_file" IS 'CA file';
