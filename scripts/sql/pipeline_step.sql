DROP TABLE IF EXISTS "pipeline_step";
CREATE TABLE "pipeline_step"
(
  "id"            SERIAL4     NOT NULL,
  "creation_time" TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "modified_time" TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "pipeline_id"   VARCHAR(21) NOT NULL,

  PRIMARY KEY ("id")
);

COMMENT ON TABLE "pipeline_step" IS 'Pipeline';

COMMENT ON COLUMN "pipeline_step"."id" IS 'PK';
COMMENT ON COLUMN "pipeline_step"."creation_time" IS 'Creation timestamp';
COMMENT ON COLUMN "pipeline_step"."modified_time" IS 'Modified timestamp';
COMMENT ON COLUMN "pipeline_step"."pipeline_id" IS 'Nano ID';
