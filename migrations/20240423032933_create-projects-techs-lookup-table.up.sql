CREATE TABLE projects_technologies (
    project_id INTEGER NOT NULL REFERENCES projects(id),
    technology_id INTEGER NOT NULL REFERENCES technologies(id),
    PRIMARY KEY (project_id, technology_id) -- ensures a project-technology pair is unique
);
