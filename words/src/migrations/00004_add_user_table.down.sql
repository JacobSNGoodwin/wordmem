ALTER TABLE words
DROP CONSTRAINT constraint_user_id_fk;

ALTER TABLE words
RENAME COLUMN user_id TO userid;

DROP TABLE users;