-- name: GetAllFacturas :many
SELECT 
    f.idFactura,
    f.fechaFact,
    f.metodoPago,
    f.iva,
    f.descuento,
    f.subtotal,
    f.total,
    r.numReserva,
    r.fechaReserva,
    r.horaReserva,
    r.cantidadPersonas,
    p.nombre AS nombreCliente,
    p.apellido_1,
    p.apellido_2
FROM reservas.Factura f
JOIN reservas.Reserva r ON f.reserva = r.numReserva
JOIN reservas.Persona p ON r.persona = p.idPersona;

-- name: GetFacturaById :one
SELECT 
    f.idFactura,
    f.fechaFact,
    f.metodoPago,
    f.iva,
    f.descuento,
    f.subtotal,
    f.total,
    r.numReserva,
    r.fechaReserva,
    r.horaReserva,
    r.cantidadPersonas,
    p.nombre AS nombreCliente,
    p.apellido_1,
    p.apellido_2
FROM reservas.Factura f
JOIN reservas.Reserva r ON f.reserva = r.numReserva
JOIN reservas.Persona p ON r.persona = p.idPersona
WHERE f.idFactura = ?;

-- name: CreateFactura :execresult
INSERT INTO reservas.Factura (idFactura, fechaFact, metodoPago, iva, descuento, subtotal, total, reserva)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: DeleteFactura :exec
DELETE FROM reservas.Factura WHERE idFactura = ?;

-- name: UpdateFactura :exec
UPDATE reservas.Factura
SET metodoPago = ?, iva = ?, descuento = ?, subtotal = ?, total = ?
WHERE idFactura = ?;