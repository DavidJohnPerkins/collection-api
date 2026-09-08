package store

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb"
)

const driverName = "sqlserver"

type SqlServerCollectionStore struct {
	dbx *sqlx.DB
}

func NewSqlServerCollectionStore(dbx *sqlx.DB) *SqlServerCollectionStore {
	return &SqlServerCollectionStore{dbx: dbx}
}

func noOpMapper(s string) string {
	return s
}

func InitSharedDB(ctx context.Context, databaseURL string) (*sqlx.DB, error) {
	dbx, err := sqlx.ConnectContext(ctx, driverName, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("connect failed: %w", err)
	}

	dbx.MapperFunc(noOpMapper)

	// Connection pool settings
	dbx.SetMaxOpenConns(25)
	dbx.SetMaxIdleConns(10)
	dbx.SetConnMaxLifetime(5 * time.Minute)
	dbx.SetConnMaxIdleTime(1 * time.Minute)

	return dbx, nil
}

func (s *SqlServerCollectionStore) close() error {
	return s.dbx.Close()
}

func (s *SqlServerCollectionStore) GetOSMapList(ctx context.Context, mapRange string) ([]OSMap, error) {

	var maps []OSMap
	sqlCmd := `EXEC COLLECTION.r_OS_` + strings.ToUpper(mapRange) + ` @p_input_json = @json`
	jsonBody := `{"item_id": -1}`

	r, err := s.dbx.QueryxContext(
		ctx,
		sqlCmd,
		sql.Named("json", jsonBody))

	if err != nil {
		return nil, err
	}
	defer r.Close()

	for r.Next() {
		var m OSMap
		if err := r.StructScan(&m); err != nil {
			log.Printf("failed: %v", err)
			return nil, err
		}
		maps = append(maps, m)
	}

	return maps, nil
}

func (s *SqlServerCollectionStore) GetOSMapItem(ctx context.Context, mapRange string, item_id int) (OSMap, error) {

	var m OSMap
	sqlCmd := `EXEC COLLECTION.r_OS_` + strings.ToUpper(mapRange) + ` @p_input_json = @json`
	jsonBody := fmt.Sprintf(`{"item_id": %d}`, item_id)

	r, err := s.dbx.QueryxContext(
		ctx,
		sqlCmd,
		sql.Named("json", jsonBody))

	if err != nil {
		return OSMap{}, err
	}
	defer r.Close()

	for r.Next() {
		if err := r.StructScan(&m); err != nil {
			log.Printf("failed: %v", err)
			return OSMap{}, err
		}
	}

	return m, nil
}

func (s *SqlServerCollectionStore) GetInkList(ctx context.Context) ([]Ink, error) {

	var inks []Ink
	sqlCmd := `EXEC COLLECTION.r_INK_COLLECTION @p_input_json = @json`
	jsonBody := `{"item_id": -1}`

	r, err := s.dbx.QueryxContext(
		ctx,
		sqlCmd,
		sql.Named("json", jsonBody))

	if err != nil {
		return nil, err
	}
	defer r.Close()

	for r.Next() {
		var i Ink
		if err := r.StructScan(&i); err != nil {
			log.Printf("failed: %v", err)
			return nil, err
		}
		inks = append(inks, i)
	}

	return inks, nil
}

func (s *SqlServerCollectionStore) GetInkItem(ctx context.Context, item_id int) (Ink, error) {

	var ink Ink
	sqlCmd := `EXEC COLLECTION.r_INK_COLLECTION @p_input_json = @json`
	jsonBody := fmt.Sprintf(`{"item_id": %d}`, item_id)

	r, err := s.dbx.QueryxContext(
		ctx,
		sqlCmd,
		sql.Named("json", jsonBody))

	if err != nil {
		return Ink{}, err
	}
	defer r.Close()

	for r.Next() {
		if err := r.StructScan(&ink); err != nil {
			log.Printf("failed: %v", err)
			return Ink{}, err
		}
	}

	return ink, nil
}

func (s *SqlServerCollectionStore) GetPenList(ctx context.Context) ([]Pen, error) {

	var pens []Pen
	sqlCmd := `EXEC COLLECTION.r_PEN_COLLECTION @p_input_json = @json`
	jsonBody := `{"item_id": -1}`

	r, err := s.dbx.QueryxContext(
		ctx,
		sqlCmd,
		sql.Named("json", jsonBody))

	if err != nil {
		return nil, err
	}
	defer r.Close()

	for r.Next() {
		var p Pen
		if err := r.StructScan(&p); err != nil {
			log.Printf("failed: %v", err)
			return nil, err
		}
		pens = append(pens, p)
	}

	return pens, nil
}

func (s *SqlServerCollectionStore) GetPenItem(ctx context.Context, item_id int) (Pen, error) {

	var pen Pen
	sqlCmd := `EXEC COLLECTION.r_PEN_COLLECTION @p_input_json = @json`
	jsonBody := fmt.Sprintf(`{"item_id": %d}`, item_id)

	r, err := s.dbx.QueryxContext(
		ctx,
		sqlCmd,
		sql.Named("json", jsonBody))

	if err != nil {
		return Pen{}, err
	}
	defer r.Close()

	for r.Next() {
		if err := r.StructScan(&pen); err != nil {
			log.Printf("failed: %v", err)
			return Pen{}, err
		}
	}

	return pen, nil
}

func (s *SqlServerCollectionStore) GetScoreList(ctx context.Context) ([]Score, error) {

	var scores []Score
	sqlCmd := `EXEC COLLECTION.r_SCORES @p_input_json = @json`
	jsonBody := `{"item_id": -1}`

	r, err := s.dbx.QueryxContext(
		ctx,
		sqlCmd,
		sql.Named("json", jsonBody))

	if err != nil {
		return nil, err
	}
	defer r.Close()

	for r.Next() {
		var s Score
		if err := r.StructScan(&s); err != nil {
			log.Printf("failed: %v", err)
			return nil, err
		}
		scores = append(scores, s)
	}

	return scores, nil
}

func (s *SqlServerCollectionStore) GetScoreItem(ctx context.Context, item_id int) (Score, error) {

	var score Score
	sqlCmd := `EXEC COLLECTION.r_SCORES @p_input_json = @json`
	jsonBody := fmt.Sprintf(`{"item_id": %d}`, item_id)

	r, err := s.dbx.QueryxContext(
		ctx,
		sqlCmd,
		sql.Named("json", jsonBody))

	if err != nil {
		return Score{}, err
	}
	defer r.Close()

	for r.Next() {
		if err := r.StructScan(&score); err != nil {
			log.Printf("failed: %v", err)
			return Score{}, err
		}
	}

	return score, nil
}

func (s *SqlServerCollectionStore) GetPolychromList(ctx context.Context) ([]Polychrom, error) {

	var pencils []Polychrom
	sqlCmd := `EXEC COLLECTION.r_POLYCHROMOS_PENCILS @p_input_json = @json`
	jsonBody := `{"item_id": -1}`

	r, err := s.dbx.QueryxContext(
		ctx,
		sqlCmd,
		sql.Named("json", jsonBody))

	if err != nil {
		return nil, err
	}
	defer r.Close()

	for r.Next() {
		var p Polychrom
		if err := r.StructScan(&p); err != nil {
			log.Printf("failed: %v", err)
			return nil, err
		}
		pencils = append(pencils, p)
	}

	return pencils, nil
}

func (s *SqlServerCollectionStore) GetPolychromItem(ctx context.Context, item_id int) (Polychrom, error) {

	var pencil Polychrom
	sqlCmd := `EXEC COLLECTION.r_POLYCHROMOS_PENCILS @p_input_json = @json`
	jsonBody := fmt.Sprintf(`{"item_id": %d}`, item_id)

	r, err := s.dbx.QueryxContext(
		ctx,
		sqlCmd,
		sql.Named("json", jsonBody))

	if err != nil {
		return Polychrom{}, err
	}
	defer r.Close()

	for r.Next() {
		if err := r.StructScan(&pencil); err != nil {
			log.Printf("failed: %v", err)
			return Polychrom{}, err
		}
	}

	return pencil, nil
}
