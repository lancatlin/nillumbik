-- name: GetSite :one
SELECT * FROM sites
WHERE id = $1 LIMIT 1;

-- name: ListSites :many
SELECT * FROM sites
ORDER BY name;