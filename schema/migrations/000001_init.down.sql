-- Drop the trigger before dropping the table
DROP TRIGGER IF EXISTS trg_update_updated_at ON tbl_vacancy;

-- Drop the foreign key constraint before dropping the table
ALTER TABLE tbl_vacancy DROP CONSTRAINT tbl_vacancy_vacancy_team_id_fkey;

-- Drop the tables
DROP TABLE tbl_vacancy;
DROP TABLE tbl_team;
