-- name: CreateUsuario :execresult
INSERT INTO reservas.Usuario (userName, password, rol, idPersona)
VALUES (?, ?, ?, ?);

-- name: GetUsuarioByUserName :one
SELECT userName, password, rol, idPersona
FROM reservas.Usuario
WHERE userName = ?;

-- name: UpdateUsuario :exec
UPDATE reservas.Usuario
SET password = ?
WHERE userName = ?;

-- name: DeleteUsuario :exec
DELETE FROM reservas.Usuario
WHERE userName = ?;

-- name: GetAllUsuarios :many
SELECT userName, rol, idPersona
FROM reservas.Usuario;

-- name: UsuarioExiste :one
SELECT COUNT(*) as count
FROM reservas.Usuario
WHERE userName = ?;


-- name: GetCorreoByUserName :one
SELECT p.correo
FROM reservas.Usuario u
JOIN reservas.Persona p ON u.idPersona = p.idPersona
WHERE u.userName = ?;

-- name: GetUsuarioByCorreo :one
SELECT u.userName, u.password, u.rol, u.idPersona
FROM reservas.Usuario u
JOIN reservas.Persona p ON u.idPersona = p.idPersona
WHERE p.correo = ?;
