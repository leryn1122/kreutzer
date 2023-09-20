DROP TABLE IF EXISTS "webhook";
CREATE TABLE "webhook"
(
  "id"            SERIAL4     NOT NULL,
  "creation_time" TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "modified_time" TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "hook_id"       VARCHAR(21) NOT NULL,
  "channel"       VARCHAR(12) NOT NULL,
  "enabled"       BOOL        NOT NULL DEFAULT TRUE,
  "url"           VARCHAR(64) NOT NULL,
  "secret"        VARCHAR(64),

  PRIMARY KEY ("id"),
  UNIQUE ("hook_id", "channel")
);

COMMENT ON TABLE "webhook" IS 'Webhooks';

COMMENT ON COLUMN "webhook"."id" IS 'PK';
COMMENT ON COLUMN "webhook"."creation_time" IS 'Creation timestamp';
COMMENT ON COLUMN "webhook"."modified_time" IS 'Modified timestamp';
COMMENT ON COLUMN "webhook"."hook_id" IS 'Webhook ID';
COMMENT ON COLUMN "webhook"."channel" IS 'SCM channel:`github`-GitHub';
COMMENT ON COLUMN "webhook"."enabled" IS 'Enabled or not';
COMMENT ON COLUMN "webhook"."url" IS 'Webhook callback URL';
COMMENT ON COLUMN "webhook"."secret" IS 'Webhook secret';
