DROP TABLE IF EXISTS "cluster";
CREATE TABLE "cluster"
(
  "id"            SERIAL4   NOT NULL,
  "creation_time" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "modified_time" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "cluster"       VARCHAR(64) NULL,
  PRIMARY KEY ("id")
);

COMMENT
ON TABLE "cluster" IS 'Cluster';

COMMENT
ON COLUMN "cluster"."id"              IS 'PK';
COMMENT
ON COLUMN "cluster"."creation_time"   IS 'Creation timestamp';
COMMENT
ON COLUMN "cluster"."modified_time"   IS 'Modified timestamp';
COMMENT
ON COLUMN "cluster"."cluster"         IS 'Cluster Info';

