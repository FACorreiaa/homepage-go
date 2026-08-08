-- name: RecordVisit :exec
INSERT INTO visits (visitor_hash, country_code, city, lat, lon, path)
VALUES (?, ?, ?, ?, ?, ?);

-- name: ListRecentVisits :many
SELECT visitor_hash, country_code, city, lat, lon, path, created_at
FROM visits
WHERE created_at >= datetime('now', ?)
ORDER BY created_at DESC
LIMIT ?;

-- name: CountVisitsSince :one
SELECT COUNT(*) FROM visits WHERE created_at >= datetime('now', ?);

-- name: CountVisitorsSince :one
SELECT COUNT(DISTINCT visitor_hash) FROM visits WHERE created_at >= datetime('now', ?);

-- name: CountCountriesSince :one
SELECT COUNT(DISTINCT country_code) FROM visits WHERE created_at >= datetime('now', ?);

-- name: TopCountriesSince :many
SELECT country_code, COUNT(*) AS visits
FROM visits
WHERE created_at >= datetime('now', ?)
GROUP BY country_code
ORDER BY visits DESC, country_code ASC
LIMIT ?;

-- name: TopPathsSince :many
SELECT path, COUNT(*) AS visits
FROM visits
WHERE created_at >= datetime('now', ?)
GROUP BY path
ORDER BY visits DESC, path ASC
LIMIT ?;

-- name: DeleteVisitsBefore :exec
DELETE FROM visits WHERE created_at < datetime('now', ?);

-- name: GetGeoCache :one
SELECT prefix, country_code, city, lat, lon, created_at
FROM geo_cache
WHERE prefix = ? LIMIT 1;

-- name: UpsertGeoCache :exec
INSERT INTO geo_cache (prefix, country_code, city, lat, lon)
VALUES (?, ?, ?, ?, ?)
ON CONFLICT(prefix) DO UPDATE SET
    country_code = excluded.country_code,
    city         = excluded.city,
    lat          = excluded.lat,
    lon          = excluded.lon;
