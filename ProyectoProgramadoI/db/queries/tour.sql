-- name: GetAllTours :many
SELECT * FROM reservas.Tour;

-- name: GetTourById :one
SELECT * FROM reservas.Tour WHERE idTour = ? LIMIT 1;

-- name: CreateTour :execresult
INSERT INTO reservas.Tour (idTour, nombre, descripcion, tipo, disponibilidad, precioBase, ubicacion)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdateTour :exec
UPDATE reservas.Tour
SET nombre = ?, descripcion = ?, tipo = ?, disponibilidad = ?, precioBase = ?, ubicacion = ?
WHERE idTour = ?;

-- name: DeleteTour :exec
DELETE FROM reservas.Tour WHERE idTour = ?;