package store

import (
	"context"
)

type OSMap struct {
	Item_id          int    `db:"item_id"`
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

type Ink struct {
	Item_id        int    `db:"item_id"`
	Brand          string `db:"BRAND"`
	Range          string `db:"RANGE"`
	Container      string `db:"CONTAINER"`
	Container_size string `db:"CONTAINER_SIZE"`
	Colour         string `db:"COLOUR"`
	Qty_remaining  int    `db:"QTY_REMAINING"`
	Rating         int    `db:"RATING"`
	Ink_name       string `db:"INK_NAME"`
	Key_Value      string `db:"key_value"`
	Image_1        string `db:"IMAGE_1"`
	Comments       string `db:"COMMENTS"`
}

type Interface interface {
	GetOSMapList(ctx context.Context, mapRange string) ([]OSMap, error)
	GetOSMapItem(ctx context.Context, mapRange string, item_id int) (OSMap, error)
	GetInkList(ctx context.Context) ([]Ink, error)
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
