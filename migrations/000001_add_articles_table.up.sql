CREATE TABLE IF NOT EXISTS articles
(
    id           bigserial PRIMARY KEY,
    title        text                        NOT NULL,
    slug         text                        NOT NULL,
    content      text                        NOT NULL,
    created_at   timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at   timestamp(0) with time zone,
    published_at timestamp(0) with time zone,
    version      integer                     NOT NULL DEFAULT 1
);

-- Modify the table to add a UNIQUE constraint on slug
ALTER TABLE articles
    ADD CONSTRAINT articles_slug_unique UNIQUE (slug);

-- Create or replace the trigger function to auto-update updated_at and increment version
CREATE OR REPLACE FUNCTION update_article_metadata()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at := NOW();
    NEW.version := OLD.version + 1;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Drop existing trigger if it exists (optional safety step)
DROP TRIGGER IF EXISTS set_article_metadata ON articles;

-- Create the trigger to fire before UPDATE
CREATE TRIGGER set_article_metadata
    BEFORE UPDATE
    ON articles
    FOR EACH ROW
EXECUTE FUNCTION update_article_metadata();