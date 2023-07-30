DROP TABLE IF EXISTS "schedule";
CREATE TABLE "schedule"
(
  "id"              SERIAL4     NOT NULL,
  "creation_time"   TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "modified_time"   TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "cron_expression" VARCHAR(64) NULL,
  PRIMARY KEY ("id")
);

COMMENT ON TABLE "schedule" IS 'Schedule';

COMMENT ON COLUMN "schedule"."id" IS 'PK';
COMMENT ON COLUMN "schedule"."creation_time" IS 'Creation timestamp';
COMMENT ON COLUMN "schedule"."modified_time" IS 'Modified timestamp';
COMMENT ON COLUMN "schedule"."cron_expression" IS 'Cron expression';
