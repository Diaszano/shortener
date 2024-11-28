CREATE TABLE IF NOT EXISTS "availability"(
    "id" BIGINT NOT NULL PRIMARY KEY,
    "shortened_id" BIGINT NOT NULL REFERENCES "shortened"("id"),
    "start" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "end" TIMESTAMPTZ NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE("shortened_id")
);
CREATE TRIGGER update_updated_at BEFORE UPDATE ON "availability" FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
COMMENT
    ON TABLE
    "availability" IS 'The availability table stores the start and end dates during which a shortened URL can be accessed. This table manages the time-bound availability of shortened links.';
COMMENT
    ON COLUMN
    "availability"."id" IS 'A unique identifier for each availability record.';
COMMENT
    ON COLUMN
    "availability"."shortened_id" IS 'The identifier of the shortened URL that this availability record is associated with. This links the availability period to a specific shortened URL.';
COMMENT
    ON COLUMN
    "availability"."start" IS 'The date and time when access to the shortened URL becomes available.';
COMMENT
    ON COLUMN
    "availability"."end" IS 'The date and time when access to the shortened URL will expire.';
COMMENT
    ON COLUMN
    "availability"."created_at" IS 'The timestamp indicating when the availability record was created.';
COMMENT
    ON COLUMN
    "availability"."updated_at" IS 'The timestamp of the most recent update made to the availability record.';
