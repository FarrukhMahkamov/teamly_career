CREATE TABLE tbl_team (
    team_id SERIAL PRIMARY KEY,
    team_name VARCHAR(255) NOT NULL
);

CREATE TABLE tbl_user (
    user_id SERIAL PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    user_second_name VARCHAR(255) NOT NULL,
    user_email VARCHAR(255) NOT NULL,
    user_phone VARCHAR(255) NOT NULL,
    user_password VARCHAR(255) NOT NULL,
    user_photo VARCHAR(255) NOT NULL,
    user_status INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE tbl_vacancy 
(
    vacancy_id SERIAL PRIMARY KEY,
    vacancy_team_id INT NOT NULL,
    vacancy_status_id INT NOT NULL,
    position VARCHAR(255) NOT NULL,
    vacancy_description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (vacancy_team_id) REFERENCES tbl_team(team_id)
);

CREATE TABLE tbl_vacancy_detail 
(
    vacancy_detail_id SERIAL PRIMARY KEY,
    vacancy_id INT NOT NULL,
    apply_count INT DEFAULT 0,
    level VARCHAR(255) NOT NULL,
    experience VARCHAR(255) NOT NULL,
    work_type VARCHAR(255) NOT NULL,
    work_time VARCHAR(255) NOT NULL,
    work_location VARCHAR(255) NOT NULL,
    salary VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (vacancy_id) REFERENCES tbl_vacancy(vacancy_id)
);

CREATE TABLE tbl_user_file (
    user_file_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    file_size INT NOT NULL,
    file_type VARCHAR(255) NOT NULL,
    status_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (user_id) REFERENCES tbl_user(user_id)
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_update_updated_at_detail
BEFORE UPDATE ON tbl_vacancy_detail
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER trg_update_updated_at
BEFORE UPDATE ON tbl_vacancy
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER trg_update_updated_at_user_file
BEFORE UPDATE ON tbl_user_file
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();



INSERT INTO tbl_team (team_name) VALUES ('Frontend');
INSERT INTO tbl_team (team_name) VALUES ('Backend');
INSERT INTO tbl_team (team_name) VALUES ('DevOps');
INSERT INTO tbl_team (team_name) VALUES ('QA');
INSERT INTO tbl_team (team_name) VALUES ('UX/UI');
INSERT INTO tbl_team (team_name) VALUES ('Product');
INSERT INTO tbl_team (team_name) VALUES ('Marketing');
INSERT INTO tbl_team (team_name) VALUES ('Sales');
INSERT INTO tbl_team (team_name) VALUES ('HR');
INSERT INTO tbl_team (team_name) VALUES ('Database');

