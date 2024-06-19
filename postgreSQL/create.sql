CREATE TABLE Users (
    userId BIGINT PRIMARY KEY NOT NULL,
    name text DEFAULT 'missed',
    lastname text DEFAULT 'missed',
    phone text DEFAULT 'missed',
    nickname text DEFAULT '',
    dvijid int DEFAULT 0,
    language text DEFAULT '',
    action text DEFAULT 'new',
    level integer DEFAULT 0,
    isadmin integer DEFAULT 0,
    titleru text DEFAULT '',
    titleen text DEFAULT '',
    discrpru text DEFAULT '',
    discrpen text DEFAULT '',
    delactivity integer DEFAULT 0,
    changeable text DEFAULT '',
    changesru text DEFAULT '',
    changesen text DEFAULT ''
);

CREATE TABLE Phrases (
    name SERIAL PRIMARY KEY,
    titleru text DEFAULT '',
    titleen text DEFAULT '',
    discrpru text DEFAULT '',
    discrpen text DEFAULT '',
    status int DEFAULT 0
);

CREATE TABLE Dvij (
    id SERIAL PRIMARY KEY,
    caption text DEFAULT '',
    description text DEFAULT '',
    datetime TIMESTAMPTZ,
    link text DEFAULT '',
    status int DEFAULT 0
);

CREATE TABLE DvijClients (
    id integer PRIMARY KEY,
    userid BIGINT DEFAULT 0,
    notiftime TIMESTAMPTZ,
    notifeone bool DEFAULT false,
    answerone INTEGER DEFAULT 0,
    notiftwo bool DEFAULT false,
    answertwo INTEGER DEFAULT 0,
    status integer DEFAULT 0
);