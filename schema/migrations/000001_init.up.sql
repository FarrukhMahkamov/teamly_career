CREATE TABLE tbl_teams (
    team_id SERIAL PRIMARY KEY,
    team_name VARCHAR(255) NOT NULL
);

INSERT INTO tbl_teams (team_name) VALUES ('Frontend');
INSERT INTO tbl_teams (team_name) VALUES ('Backend');
INSERT INTO tbl_teams (team_name) VALUES ('DevOps');
INSERT INTO tbl_teams (team_name) VALUES ('QA');
INSERT INTO tbl_teams (team_name) VALUES ('UX/UI');
INSERT INTO tbl_teams (team_name) VALUES ('Product');
INSERT INTO tbl_teams (team_name) VALUES ('Marketing');
INSERT INTO tbl_teams (team_name) VALUES ('Sales');
INSERT INTO tbl_teams (team_name) VALUES ('HR');