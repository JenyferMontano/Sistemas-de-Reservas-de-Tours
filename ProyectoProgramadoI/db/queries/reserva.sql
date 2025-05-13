-- name: CreateReserva :execresult  
INSERT INTO reservas.Reserva (
    numReserva, fechaReserva, horaReserva, cantidadPersonas, tour, usuario, persona, transfer
) VALUES (?, ?, ?, ?, ?, ?, ?, ?);
-- name: UpdateReserva :exec
UPDATE reservas.Reserva
SET fechaReserva = ?, horaReserva = ?, cantidadPersonas = ?, tour = ?, usuario = ?, persona = ?, transfer = ?
WHERE numReserva = ?;

-- name: DeleteReserva :exec
DELETE FROM reservas.Reserva WHERE numReserva = ?;

-- name: GetReservaById :one
SELECT numReserva, fechaReserva, horaReserva, cantidadPersonas, tour, usuario, persona, transfer
FROM reservas.Reserva
WHERE numReserva = ?;

-- name: GetAllReservas :many
SELECT r.numReserva, r.fechaReserva, r.horaReserva, 
       r.cantidadPersonas, 
       t.nombre AS nombreTour, 
       u.userName AS nombreUsuario, 
       p.nombre AS nombrePersona, 
       tr.idTransfer AS idTransfer
FROM reservas.Reserva r
JOIN reservas.Tour t ON r.tour = t.idTour
JOIN reservas.Usuario u ON r.usuario = u.userName
JOIN reservas.Persona p ON r.persona = p.idPersona
JOIN reservas.Transfer tr ON r.transfer = tr.idTransfer;

-- name: GetFacturaByReservaId :one
SELECT idFactura 
FROM reservas.Factura 
WHERE reserva = ?;