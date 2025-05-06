-- name: GetAllPersonas :many
SELECT * FROM reservas.Persona;

-- name: GetPersonaById :one
SELECT * FROM reservas.Persona WHERE idPersona = ? LIMIT 1;

-- name: CreatePersona :execresult
INSERT INTO reservas.Persona (idPersona, nombre, apellido_1, apellido_2, fechaNac, telefono, correo)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdatePersona :exec
UPDATE reservas.Persona
SET nombre = ?, apellido_1 = ?, apellido_2 = ?, fechaNac = ?, telefono = ?, correo = ?
WHERE idPersona = ?;

-- name: DeletePersona :exec
DELETE FROM reservas.Persona WHERE idPersona = ?;
