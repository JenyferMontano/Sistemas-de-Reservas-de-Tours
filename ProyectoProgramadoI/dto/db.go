// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package dto

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createPersonaStmt, err = db.PrepareContext(ctx, createPersona); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePersona: %w", err)
	}
	if q.createTourStmt, err = db.PrepareContext(ctx, createTour); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTour: %w", err)
	}
	if q.deletePersonaStmt, err = db.PrepareContext(ctx, deletePersona); err != nil {
		return nil, fmt.Errorf("error preparing query DeletePersona: %w", err)
	}
	if q.deleteTourStmt, err = db.PrepareContext(ctx, deleteTour); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTour: %w", err)
	}
	if q.getAllPersonasStmt, err = db.PrepareContext(ctx, getAllPersonas); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllPersonas: %w", err)
	}
	if q.getAllToursStmt, err = db.PrepareContext(ctx, getAllTours); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllTours: %w", err)
	}
	if q.getPersonaByIdStmt, err = db.PrepareContext(ctx, getPersonaById); err != nil {
		return nil, fmt.Errorf("error preparing query GetPersonaById: %w", err)
	}
	if q.getTourByIdStmt, err = db.PrepareContext(ctx, getTourById); err != nil {
		return nil, fmt.Errorf("error preparing query GetTourById: %w", err)
	}
	if q.updatePersonaStmt, err = db.PrepareContext(ctx, updatePersona); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePersona: %w", err)
	}
	if q.updateTourStmt, err = db.PrepareContext(ctx, updateTour); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTour: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createPersonaStmt != nil {
		if cerr := q.createPersonaStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPersonaStmt: %w", cerr)
		}
	}
	if q.createTourStmt != nil {
		if cerr := q.createTourStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTourStmt: %w", cerr)
		}
	}
	if q.deletePersonaStmt != nil {
		if cerr := q.deletePersonaStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deletePersonaStmt: %w", cerr)
		}
	}
	if q.deleteTourStmt != nil {
		if cerr := q.deleteTourStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTourStmt: %w", cerr)
		}
	}
	if q.getAllPersonasStmt != nil {
		if cerr := q.getAllPersonasStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllPersonasStmt: %w", cerr)
		}
	}
	if q.getAllToursStmt != nil {
		if cerr := q.getAllToursStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllToursStmt: %w", cerr)
		}
	}
	if q.getPersonaByIdStmt != nil {
		if cerr := q.getPersonaByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPersonaByIdStmt: %w", cerr)
		}
	}
	if q.getTourByIdStmt != nil {
		if cerr := q.getTourByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTourByIdStmt: %w", cerr)
		}
	}
	if q.updatePersonaStmt != nil {
		if cerr := q.updatePersonaStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePersonaStmt: %w", cerr)
		}
	}
	if q.updateTourStmt != nil {
		if cerr := q.updateTourStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTourStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                 DBTX
	tx                 *sql.Tx
	createPersonaStmt  *sql.Stmt
	createTourStmt     *sql.Stmt
	deletePersonaStmt  *sql.Stmt
	deleteTourStmt     *sql.Stmt
	getAllPersonasStmt *sql.Stmt
	getAllToursStmt    *sql.Stmt
	getPersonaByIdStmt *sql.Stmt
	getTourByIdStmt    *sql.Stmt
	updatePersonaStmt  *sql.Stmt
	updateTourStmt     *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                 tx,
		tx:                 tx,
		createPersonaStmt:  q.createPersonaStmt,
		createTourStmt:     q.createTourStmt,
		deletePersonaStmt:  q.deletePersonaStmt,
		deleteTourStmt:     q.deleteTourStmt,
		getAllPersonasStmt: q.getAllPersonasStmt,
		getAllToursStmt:    q.getAllToursStmt,
		getPersonaByIdStmt: q.getPersonaByIdStmt,
		getTourByIdStmt:    q.getTourByIdStmt,
		updatePersonaStmt:  q.updatePersonaStmt,
		updateTourStmt:     q.updateTourStmt,
	}
}
