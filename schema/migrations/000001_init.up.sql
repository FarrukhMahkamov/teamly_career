CREATE TABLE tbl_team (
    team_id SERIAL PRIMARY KEY,
    team_name VARCHAR(255) NOT NULL
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
    apply_count INT NOT NULL,
    level VARCHAR(255) NOT NULL,
    experience VARCHAR(255) NOT NULL,
    work_type VARCHAR(255) NOT NULL,
    work_time VARCHAR(255) NOT NULL,
    work_location VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (vacancy_id) REFERENCES tbl_vacancy(vacancy_id)
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

INSERT INTO tbl_vacancy (vacancy_team_id, vacancy_status_id, position, vacancy_description) 
VALUES (1, 1, 'Frontend Developer', 'We are looking for an experienced frontend developer to join our team and help build beautiful and functional websites.');

INSERT INTO tbl_vacancy (vacancy_team_id, vacancy_status_id, position, vacancy_description) 
VALUES (2, 1, 'Backend Developer', 'We are seeking a talented backend developer to help us build and maintain our scalable and efficient server-side applications.');

INSERT INTO tbl_vacancy (vacancy_team_id, vacancy_status_id, position, vacancy_description) 
VALUES (3, 1, 'DevOps Engineer', 'We are looking for a DevOps engineer to help us build and maintain our infrastructure and deployment pipelines.');

INSERT INTO tbl_vacancy (vacancy_team_id, vacancy_status_id, position, vacancy_description) 
VALUES (4, 1, 'QA Tester', 'We are seeking a skilled QA tester to help us ensure the quality of our products and applications.');

INSERT INTO tbl_vacancy (vacancy_team_id, vacancy_status_id, position, vacancy_description) 
VALUES (5, 1, 'UX/UI Designer', 'We are seeking a talented UX/UI designer to help us create beautiful and intuitive user interfaces for our products and applications.');

INSERT INTO tbl_vacancy (vacancy_team_id, vacancy_status_id, position, vacancy_description) 
VALUES (6, 1, 'Product Manager', 'We are seeking a passionate and experienced product manager to help us develop and launch new products and features.');

INSERT INTO tbl_vacancy (vacancy_team_id, vacancy_status_id, position, vacancy_description) 
VALUES (7, 1, 'Marketing Manager', 'We are seeking a skilled marketing manager to help us develop and implement effective marketing strategies and campaigns.');

INSERT INTO tbl_vacancy (vacancy_team_id, vacancy_status_id, position, vacancy_description) 
VALUES (8, 1, 'Sales Representative', 'We are seeking a motivated and experienced sales representative to help us acquire new customers and increase revenue.');

INSERT INTO tbl_vacancy (vacancy_team_id, vacancy_status_id, position, vacancy_description) 
VALUES (9, 1, 'HR Manager', 'We are seeking an experienced HR manager to help us attract, retain, and develop top talent.');

INSERT INTO tbl_vacancy (vacancy_team_id, vacancy_status_id, position, vacancy_description) 
VALUES (10, 1, 'Database Administrator', 'We are seeking a skilled database administrator to help us maintain and optimize our databases and data pipelines.');
