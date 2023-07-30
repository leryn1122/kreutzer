CREATE OR REPLACE FUNCTION "fn_update_at_modified_time"()
  RETURNS TRIGGER
  LANGUAGE "plpgsql"
AS
$function$
BEGIN
  "new"."modified_time" = NOW();
  RETURN "new";
END;
$function$
;

CREATE OR REPLACE FUNCTION "fn_create_update_trigger_on_ddl"()
  RETURNS EVENT_TRIGGER
  LANGUAGE "plpgsql"
  SECURITY INVOKER
AS
$function$
DECLARE
  "r"                  RECORD;
  DECLARE "sql"        TEXT;
  DECLARE "table_name" TEXT;
BEGIN
  FOR "r" IN SELECT * FROM "pg_event_trigger_ddl_commands"()
    LOOP
      IF "r"."command_tag" = 'CREATE TABLE' AND "r"."object_type" = 'table' THEN
        "table_name" = SPLIT_PART("r"."object_identity", '.', 2);
        "sql" := 'CREATE TRIGGER "trg_' || "table_name" || '_at_modification"'
                   || '  BEFORE '
                   || '    UPDATE '
                   || '  ON '
                   || '    "' || "table_name" || '" '
                   || '  FOR EACH ROW '
          || 'EXECUTE FUNCTION "fn_update_at_modified_time"() ';
        EXECUTE "sql";
      END IF;
    END LOOP;
END
$function$
;

DROP EVENT TRIGGER IF EXISTS "tg_on_create_ddl";
CREATE EVENT TRIGGER "tg_on_create_ddl" ON DDL_COMMAND_END
EXECUTE FUNCTION "fn_create_update_trigger_on_ddl"();
