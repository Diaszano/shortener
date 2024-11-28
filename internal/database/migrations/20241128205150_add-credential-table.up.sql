CREATE TABLE IF NOT EXISTS "credential"(
    "id" BIGINT NOT NULL PRIMARY KEY,
    "shortened_id" BIGINT NOT NULL REFERENCES "shortened"("id"),
    "password" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE("shortened_id")
);
CREATE TRIGGER update_updated_at BEFORE UPDATE ON "credential" FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
COMMENT
    ON TABLE
    "credential" IS 'The credential table stores all access credentials for shortened URLs. It allows for password protection of specific shortened links, ensuring that only authorized users can access them.';
COMMENT
    ON COLUMN
    "credential"."id" IS 'A unique identifier for each credential record.';
COMMENT
    ON COLUMN
    "credential"."shortened_id" IS 'The identifier of the shortened URL that this credential is associated with. This links the credential to a specific shortened URL.';
COMMENT
    ON COLUMN
    "credential"."password" IS 'The password required to access the shortened URL.';
COMMENT
    ON COLUMN
    "credential"."created_at" IS 'The timestamp indicating when the credential record was created.';
COMMENT
    ON COLUMN
    "credential"."updated_at" IS 'The timestamp of the most recent update made to the credential record.';
