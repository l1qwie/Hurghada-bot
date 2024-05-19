CREATE TABLE Users (
    userId BIGINT PRIMARY KEY NOT NULL,
    language text DEFAULT '',
    action text DEFAULT 'new',
    level integer DEFAULT 0,
    isadmin integer DEFAULT 0
)