// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: persona.sql

package dto

import (
	"context"
	"database/sql"
	"time"
)

const createPersona = `-- name: CreatePersona :execresult
INSERT INTO reservas.Persona (idPersona, nombre, apellido_1, apellido_2, fechaNac, telefono, correo)
VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreatePersonaParams struct {
	Idpersona int32     `json:"idpersona"`
	Nombre    string    `json:"nombre"`
	Apellido1 string    `json:"apellido_1"`
	Apellido2 string    `json:"apellido_2"`
	Fechanac  time.Time `json:"fechanac"`
	Telefono  string    `json:"telefono"`
	Correo    string    `json:"correo"`
}

func (q *Queries) CreatePersona(ctx context.Context, arg CreatePersonaParams) (sql.Result, error) {
	return q.exec(ctx, q.createPersonaStmt, createPersona,
		arg.Idpersona,
		arg.Nombre,
		arg.Apellido1,
		arg.Apellido2,
		arg.Fechanac,
		arg.Telefono,
		arg.Correo,
	)
}

const deletePersona = `-- name: DeletePersona :exec
DELETE FROM reservas.Persona WHERE idPersona = ?
`

func (q *Queries) DeletePersona(ctx context.Context, idpersona int32) error {
	_, err := q.exec(ctx, q.deletePersonaStmt, deletePersona, idpersona)
	return err
}

const getAllPersonas = `-- name: GetAllPersonas :many
SELECT idpersona, nombre, apellido_1, apellido_2, fechanac, telefono, correo FROM reservas.Persona
`

func (q *Queries) GetAllPersonas(ctx context.Context) ([]ReservasPersona, error) {
	rows, err := q.query(ctx, q.getAllPersonasStmt, getAllPersonas)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ReservasPersona
	for rows.Next() {
		var i ReservasPersona
		if err := rows.Scan(
			&i.Idpersona,
			&i.Nombre,
			&i.Apellido1,
			&i.Apellido2,
			&i.Fechanac,
			&i.Telefono,
			&i.Correo,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPersonaById = `-- name: GetPersonaById :one
SELECT idpersona, nombre, apellido_1, apellido_2, fechanac, telefono, correo FROM reservas.Persona WHERE idPersona = ? LIMIT 1
`

func (q *Queries) GetPersonaById(ctx context.Context, idpersona int32) (ReservasPersona, error) {
	row := q.queryRow(ctx, q.getPersonaByIdStmt, getPersonaById, idpersona)
	var i ReservasPersona
	err := row.Scan(
		&i.Idpersona,
		&i.Nombre,
		&i.Apellido1,
		&i.Apellido2,
		&i.Fechanac,
		&i.Telefono,
		&i.Correo,
	)
	return i, err
}

const updatePersona = `-- name: UpdatePersona :exec
UPDATE reservas.Persona
SET nombre = ?, apellido_1 = ?, apellido_2 = ?, fechaNac = ?, telefono = ?, correo = ?
WHERE idPersona = ?
`

type UpdatePersonaParams struct {
	Nombre    string    `json:"nombre"`
	Apellido1 string    `json:"apellido_1"`
	Apellido2 string    `json:"apellido_2"`
	Fechanac  time.Time `json:"fechanac"`
	Telefono  string    `json:"telefono"`
	Correo    string    `json:"correo"`
	Idpersona int32     `json:"idpersona"`
}

func (q *Queries) UpdatePersona(ctx context.Context, arg UpdatePersonaParams) error {
	_, err := q.exec(ctx, q.updatePersonaStmt, updatePersona,
		arg.Nombre,
		arg.Apellido1,
		arg.Apellido2,
		arg.Fechanac,
		arg.Telefono,
		arg.Correo,
		arg.Idpersona,
	)
	return err
}
