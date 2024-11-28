-- 20241128025932_add-shortened-table.up.sql
CREATE TABLE IF NOT EXISTS shortened(
    "id" BIGINT NOT NULL PRIMARY KEY,
    "path" VARCHAR(50) NOT NULL,
    "destination" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL,

    unique("path")
);
COMMENT
    ON TABLE
    shortened IS 'The shortened table stores all destination URLs and their corresponding shortened paths. This table serves as the core for managing URL shortening functionality.';
COMMENT
    ON COLUMN
    shortened."id" IS 'A unique identifier for each shortened URL record.';
COMMENT
    ON COLUMN
    shortened."path" IS 'The shortened path of the URL, representing the user-friendly and compact version.';
COMMENT
    ON COLUMN
    shortened."destination" IS 'The original, non-shortened URL that the path redirects to.';
COMMENT
    ON COLUMN
    shortened."created_at" IS 'The timestamp indicating when the shortened URL record was created.';
COMMENT
    ON COLUMN
    shortened."updated_at" IS 'The timestamp of the most recent update made to the shortened URL record.';