package store

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb"
)

const driverName = "sqlserver"

type SqlServerCollectionStore struct {
	databaseUrl string
	dbx         *sqlx.DB
}

func NewSqlServerCollectionStore(databaseUrl string) *SqlServerCollectionStore {
	return &SqlServerCollectionStore{
		databaseUrl: databaseUrl,
	}
}

func noOpMapper(s string) string {
	return s
}

func (s *SqlServerCollectionStore) connect(ctx context.Context) error {
	dbx, err := sqlx.ConnectContext(ctx, driverName, s.databaseUrl)
	if err != nil {
		log.Printf("DB connect failed: %v", err)
		return err
	}

	dbx.MapperFunc(noOpMapper)
	s.dbx = dbx
	return nil
}

func (s *SqlServerCollectionStore) close() error {
	return s.dbx.Close()
}

func (s *SqlServerCollectionStore) GetOSMapList(ctx context.Context, mapRange string) ([]OSMap, error) {
	err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer s.close()

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
	err := s.connect(ctx)
	if err != nil {
		return OSMap{}, err
	}
	defer s.close()

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
	err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer s.close()

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
	err := s.connect(ctx)
	if err != nil {
		return Ink{}, err
	}
	defer s.close()

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
	err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer s.close()

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
	err := s.connect(ctx)
	if err != nil {
		return Pen{}, err
	}
	defer s.close()

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
