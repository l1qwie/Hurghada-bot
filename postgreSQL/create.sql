CREATE TABLE Users (
    userId BIGINT PRIMARY KEY NOT NULL,
    language text DEFAULT '',
    action text DEFAULT 'new',
    level integer DEFAULT 0,
    isadmin integer DEFAULT 0
);

CREATE TABLE Phrases (
    name text DEFAULT '' PRIMARY KEY NOT NULL,
    en text DEFAULT '',
    ru text DEFAULT '',
    status int DEFAULT 0
);

CREATE TABLE Buttons (
    name text DEFAULT '' PRIMARY KEY NOT NULL,
    en text DEFAULT '',
    ru text DEFAULT '',
    status integer DEFAULT 0
)