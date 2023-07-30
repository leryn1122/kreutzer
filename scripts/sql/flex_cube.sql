DROP TABLE IF EXISTS "flex_cube";
CREATE TABLE "flex_cube"
(
  "id"            SERIAL4      NOT NULL,
  "creation_time" TIMESTAMP    NULL DEFAULT CURRENT_TIMESTAMP,
  "modified_time" TIMESTAMP    NULL DEFAULT CURRENT_TIMESTAMP,
  "domain"        VARCHAR(32)  NOT NULL,
  "entry_key"     VARCHAR(32)  NOT NULL,
  "entry_value"   VARCHAR(256) NULL,
  PRIMARY KEY ("id")
);

COMMENT ON TABLE "flex_cube" IS 'Flexible Cube';

COMMENT ON COLUMN "flex_cube"."id" IS 'PK';
COMMENT ON COLUMN "flex_cube"."creation_time" IS 'Creation timestamp';
COMMENT ON COLUMN "flex_cube"."modified_time" IS 'Modified timestamp';
COMMENT ON COLUMN "flex_cube"."domain" IS 'Domain';
COMMENT ON COLUMN "flex_cube"."entry_key" IS 'Entry key';
COMMENT ON COLUMN "flex_cube"."entry_value" IS 'Entry value';
