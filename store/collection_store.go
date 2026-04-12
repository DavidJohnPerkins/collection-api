package store

import (
	"context"
)

type OSMap struct {
	Item_id          string `db:"item_id"`
	Map_number       string `db:"map_number"`
	Map_title        string `db:"map_title"`
	Publish_date     string `db:"publish_date"`
	Main_settlements string `db:"main_settlements"`
	Key_value        string `db:"key_value"`
	Map_image        string `db:"map_image"`
	Map_image_rear   string `db:"map_image_rear"`
	Map_image_area   string `db:"map_image_area"`
	Comments         string `db:"comments"`
}

type Interface interface {
	GetOSMapList(ctx context.Context, mapRange string) ([]OSMap, error)
	GetOSMapItem(ctx context.Context, mapRange string, item_id int) (OSMap, error)
	//GetModel(ctx context.Context, id int) (ModelExtended, error)
	//GetMovieList(ctx context.Context) ([]Movie, error)
	//GetAttrDescList(ctx context.Context, attr_abbrev string) ([]AttrDesc, error)
	//GetFlagList(ctx context.Context, flag_type string) ([]Flag, error)
}

// type Interface interface {
// 	GetModelList(ctx context.Context) ([]Model, error)
// 	GetByID(ctx context.Context, id uuid.UUID) (Movie, error)
// 	Create(ctx context.Context, jsonString string) error
// 	Update(ctx context.Context, jsonString string) error
// 	Delete(ctx context.Context, jsonString string) error
// }
