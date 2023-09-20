DROP TABLE IF EXISTS "pipeline";
CREATE TABLE "pipeline"
(
  "id"            SERIAL4      NOT NULL,
  "creation_time" TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "modified_time" TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "pipeline_id"   VARCHAR(21)  NOT NULL,
  "name"          VARCHAR(32)  NOT NULL,
  "repo"          VARCHAR(128) NOT NULL,
  "enabled"       BOOL         NOT NULL DEFAULT TRUE,

  PRIMARY KEY ("id")
);

COMMENT ON TABLE "pipeline" IS 'Pipeline';

COMMENT ON COLUMN "pipeline"."id" IS 'PK';
COMMENT ON COLUMN "pipeline"."creation_time" IS 'Creation timestamp';
COMMENT ON COLUMN "pipeline"."modified_time" IS 'Modified timestamp';
COMMENT ON COLUMN "pipeline"."pipeline_id" IS 'Nano ID';
COMMENT ON COLUMN "pipeline"."name" IS 'Pipeline name';
COMMENT ON COLUMN "pipeline"."repo" IS 'URL or VCS adapter';
COMMENT ON COLUMN "pipeline"."enabled" IS 'Enabled or not';
