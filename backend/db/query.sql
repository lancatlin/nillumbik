-- name: GetSite :one
SELECT * FROM sites
WHERE code = $1 LIMIT 1;

-- name: ListSites :many
SELECT * FROM sites
ORDER BY name;