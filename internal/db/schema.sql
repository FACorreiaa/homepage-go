CREATE TABLE IF NOT EXISTS client_users (
    id             TEXT PRIMARY KEY,
    name           TEXT NOT NULL,
    email          TEXT NOT NULL UNIQUE,
    password_hash  TEXT NOT NULL,
    project_status TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS proposal_leads (
    id           TEXT PRIMARY KEY,
    name         TEXT NOT NULL,
    email        TEXT NOT NULL,
    company      TEXT,
    project_type TEXT NOT NULL,
    budget       TEXT,
    timeline     TEXT,
    details      TEXT NOT NULL,
    created_at   TEXT NOT NULL DEFAULT (datetime('now'))
);

CREATE TABLE IF NOT EXISTS blog_posts (
    slug       TEXT PRIMARY KEY,
    title      TEXT NOT NULL,
    summary    TEXT NOT NULL,
    category   TEXT NOT NULL,
    date       TEXT NOT NULL,
    body       TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at TEXT NOT NULL DEFAULT (datetime('now'))
);

-- Page views plotted on /stats. Deliberately holds no personal data: the
-- visitor is a salted hash that is unlinkable once the salt rotates, and the
-- coordinates are coarsened to roughly half a degree before they get here.
-- Rows are swept after 24 hours, so this table never grows without bound.
CREATE TABLE IF NOT EXISTS visits (
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    visitor_hash TEXT NOT NULL,
    country_code TEXT NOT NULL,
    city         TEXT NOT NULL,
    lat          REAL NOT NULL,
    lon          REAL NOT NULL,
    path         TEXT NOT NULL,
    created_at   TEXT NOT NULL DEFAULT (datetime('now'))
);

CREATE INDEX IF NOT EXISTS idx_visits_created_at ON visits(created_at);

-- Geo lookups keyed by network prefix (/24 for v4, /48 for v6) rather than by
-- address, so a whole subnet costs one call to the upstream provider and no
-- single address is ever stored.
CREATE TABLE IF NOT EXISTS geo_cache (
    prefix       TEXT PRIMARY KEY,
    country_code TEXT NOT NULL,
    city         TEXT NOT NULL,
    lat          REAL NOT NULL,
    lon          REAL NOT NULL,
    created_at   TEXT NOT NULL DEFAULT (datetime('now'))
);
