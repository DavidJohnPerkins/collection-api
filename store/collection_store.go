package store

import (
	"context"
)

type OSMap struct {
	Id               int    `db:"id"`
	Is_excluded      bool   `db:"is_excluded"`
	Sobriquet        string `db:"sobriquet"`
	Principal_name   string `db:"principal_name"`
	Hotness_quotient int    `db:"hotness_quotient"`
	Nationality      string `db:"nationality"`
	Ranking          string `db:"ranking"`
	Flags            string `db:"flags"`
	TH_url           string `db:"TH_url"`
	Movie_count      int    `db:"movie_count"`
}

type Interface interface {
	GetModelList(ctx context.Context, term string) ([]Model, error)
	GetModel(ctx context.Context, id int) (ModelExtended, error)
	GetMovieList(ctx context.Context) ([]Movie, error)
	GetAttrDescList(ctx context.Context, attr_abbrev string) ([]AttrDesc, error)
	GetFlagList(ctx context.Context, flag_type string) ([]Flag, error)
}

// type Interface interface {
// 	GetModelList(ctx context.Context) ([]Model, error)
// 	GetByID(ctx context.Context, id uuid.UUID) (Movie, error)
// 	Create(ctx context.Context, jsonString string) error
// 	Update(ctx context.Context, jsonString string) error
// 	Delete(ctx context.Context, jsonString string) error
// }
