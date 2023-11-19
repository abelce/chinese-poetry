{{$tableName := lowerCase .Name}}
DROP TABLE IF EXISTS "public"."tbl_{{$tableName}}";
CREATE TABLE "public"."tbl_{{$tableName}}" (
	"id" uuid NOT NULL,
	"data" jsonb NOT NULL
)
WITH (OIDS=FALSE);
ALTER TABLE "public"."tbl_{{$tableName}}" OWNER TO "postgres";

-- ----------------------------
--  Primary key structure for table tbl_{{$tableName}}
-- ----------------------------
ALTER TABLE "public"."tbl_{{$tableName}}" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;