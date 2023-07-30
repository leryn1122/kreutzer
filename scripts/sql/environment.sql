DROP TABLE IF EXISTS "environment";
CREATE TABLE "environment"
(
  "id"            SERIAL4   NOT NULL,
  "creation_time" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "modified_time" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "name"          VARCHAR   NOT NULL,
  "key"           VARCHAR   NOT NULL,
  "description"   VARCHAR   NOT NULL
);

COMMENT ON TABLE "environment" IS 'Environment Info';

COMMENT ON COLUMN "environment"."id" IS 'PK';
COMMENT ON COLUMN "environment"."creation_time" IS 'Creation timestamp';
COMMENT ON COLUMN "environment"."modified_time" IS 'Modified timestamp';
COMMENT ON COLUMN "environment"."name" IS 'Name';
COMMENT ON COLUMN "environment"."key" IS 'Key';
COMMENT ON COLUMN "environment"."description" IS 'Description';

