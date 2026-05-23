-- name: ListPaths :many
SELECT *
FROM paths
;

-- name: InsertPath :exec
INSERT INTO paths (path, information)
VALUES(?, ?)
;