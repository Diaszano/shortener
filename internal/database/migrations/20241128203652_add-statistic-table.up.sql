CREATE TABLE IF NOT EXISTS "statistic"(
    "id" BIGINT NOT NULL PRIMARY KEY,
    "shortened_id" BIGINT NOT NULL REFERENCES "shortened"("id"),
    "reference" VARCHAR(20),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX "statistic_shortened_id_index" ON "statistic"("shortened_id");
CREATE TRIGGER update_updated_at BEFORE UPDATE ON "statistic" FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
COMMENT
    ON TABLE
    "statistic" IS 'The statistic table stores all access statistics related to the shortened URLs. It tracks how and where each shortened URL is being accessed.';
COMMENT
    ON COLUMN
    "statistic"."id" IS 'A unique identifier for each statistics record.';
COMMENT
    ON COLUMN
    "statistic"."shortened_id" IS 'The identifier of the shortened URL that this statistic is associated with. This links the statistic to a specific shortened URL.';
COMMENT
    ON COLUMN
    "statistic"."reference" IS 'The reference or source from which the shortened URL is being accessed (e.g., the domain or application where the URL was clicked).';
COMMENT
    ON COLUMN
    "statistic"."created_at" IS 'The timestamp indicating when the statistics record was created.';
COMMENT
    ON COLUMN
    "statistic"."updated_at" IS 'The timestamp of the most recent update made to the statistics record.';
