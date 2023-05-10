-- Create tbl_teams table
CREATE TABLE tbl_teams (
    team_id SERIAL PRIMARY KEY,
    team_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL
);

-- Create trigger function to update the updated_at column
CREATE OR REPLACE FUNCTION update_team_updated_at_trigger_fn()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.* IS DISTINCT FROM OLD.* THEN
        NEW.updated_at = CURRENT_TIMESTAMP;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger to call the trigger function on each update
CREATE TRIGGER trigger_update_team_updated_at
BEFORE UPDATE ON tbl_teams
FOR EACH ROW
EXECUTE FUNCTION update_team_updated_at_trigger_fn();

