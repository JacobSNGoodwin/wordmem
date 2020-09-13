ALTER TABLE IF EXISTS words
ADD COLUMN IF NOT EXISTS definition VARCHAR DEFAULT 'You know the meaning, right?';