-- name: GetAllTransfers :many
SELECT * FROM reservas.Transfer;

-- name: GetTransferById :one
SELECT * FROM reservas.Transfer WHERE idTransfer = ? LIMIT 1;

-- name: CreateTransfer :execresult
INSERT INTO reservas.Transfer (idTransfer, tipo, capacidad)
VALUES (?, ?, ?);

-- name: UpdateTransfer :exec
UPDATE reservas.Transfer
SET tipo = ?, capacidad = ?
WHERE idTransfer = ?;

-- name: DeleteTransfer :exec
DELETE FROM reservas.Transfer WHERE idTransfer = ?;
