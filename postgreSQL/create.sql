CREATE TABLE Users (
    userId BIGINT PRIMARY KEY NOT NULL,
    language text DEFAULT '',
    action text DEFAULT 'new',
    level integer DEFAULT 0,
    isadmin integer DEFAULT 0,
    titleru text DEFAULT '',
    titleen text DEFAULT '',
    discrpru text DEFAULT '',
    discrpen text DEFAULT '',
    delactivity integer DEFAULT 0
);

CREATE TABLE Phrases (
    name SERIAL PRIMARY KEY,
    titleru text DEFAULT '',
    titleen text DEFAULT '',
    discrpru text DEFAULT '',
    discrpen text DEFAULT '',
    status int DEFAULT 0
);