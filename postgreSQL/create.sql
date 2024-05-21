CREATE TABLE Users (
    userId BIGINT PRIMARY KEY NOT NULL,
    language text DEFAULT '',
    action text DEFAULT 'new',
    level integer DEFAULT 0,
    isadmin integer DEFAULT 0,
    titleru text DEFAULT '',
    titleen text DEFAULT '',
    discrpru text DEFAULT '',
    discrpen text DEFAULT ''
);

CREATE TABLE Phrases (
    name INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    en text DEFAULT '',
    ru text DEFAULT '',
    status int DEFAULT 0
);

CREATE TABLE Buttons (
    name INTEGER PRIMARY KEY REFERENCES Phrases(name) NOT NULL,
    en text DEFAULT '',
    ru text DEFAULT '',
    status integer DEFAULT 0
);