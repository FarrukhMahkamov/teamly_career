-- Drop the trigger before dropping the table
DROP TRIGGER IF EXISTS trg_update_updated_at ON tbl_vacancy;
DROP TRIGGER IF EXISTS trg_update_updated_at_detail ON tbl_vacancy_detail;

-- Drop the foreign key constraint before dropping the table
ALTER TABLE tbl_vacancy DROP CONSTRAINT tbl_vacancy_vacancy_team_id_fkey;
ALTER TABLE tbl_vacancy_detail DROP CONSTRAINT tbl_vacancy_detail_vacancy_id_fkey;

-- Drop the tables
DROP TABLE tbl_vacancy;
DROP TABLE tbl_team;
DROP TABLE tbl_vacancy_detail;
DROP TABLE tbl_user;
