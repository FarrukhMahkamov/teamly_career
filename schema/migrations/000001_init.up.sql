CREATE TABLE tbl_team (
    team_id SERIAL PRIMARY KEY,
    team_name VARCHAR(255) NOT NULL
);

CREATE TABLE tbl_vacancy 
(
    vacancy_id SERIAL PRIMARY KEY,
    vacancy_team_id INT NOT NULL,
    vacany_status_id INT NOT NULL,
    position VARCHAR(255) NOT NULL,
    vacancy_description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    vacancy_name VARCHAR(255) NOT NULL,
    FOREIGN KEY (vacancy_team_id) REFERENCES tbl_team(team_id)
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

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