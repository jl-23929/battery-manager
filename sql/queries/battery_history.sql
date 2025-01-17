-- name: GetBatteryHistory :one
SELECT *
FROM battery_history
WHERE id = ?
LIMIT 1;

-- name: ListBatteryHistory :many
SELECT *
FROM battery_history
where battery_id = ?
ORDER BY timestamp desc;