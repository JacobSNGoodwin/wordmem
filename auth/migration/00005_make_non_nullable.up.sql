UPDATE users
SET name = COALESCE(name, ''),
  image_url = COALESCE(image_url, ''),
  website = COALESCE(website, '');
ALTER TABLE users
ALTER COLUMN name
SET NOT NULL,
  ALTER COLUMN name
SET DEFAULT '',
  ALTER COLUMN image_url
SET NOT NULL,
  ALTER COLUMN image_url
SET DEFAULT '',
  ALTER COLUMN website
SET NOT NULL,
  ALTER COLUMN website
SET DEFAULT '';