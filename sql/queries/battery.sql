-- name: GetBattery :one
SELECT * FROM battery
WHERE battery_id = ? LIMIT 1;

-- name: ListBatteries :many
SELECT * FROM battery
ORDER BY battery_id;