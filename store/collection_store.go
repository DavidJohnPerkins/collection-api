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
	Comments       string `db:"comments"`
}

type Pen struct {
	Item_id           int     `db:"item_id"`
	Brand             string  `db:"BRAND"`
	Range             string  `db:"RANGE"`
	Rating            int     `db:"RATING"`
	Key_Value         string  `db:"key_value"`
	Model_name        string  `db:"MODEL_NAME"`
	Body_material     string  `db:"BODY_MATERIAL"`
	Body_colour       string  `db:"BODY_COLOUR"`
	Grip_material     string  `db:"GRIP_MATERIAL"`
	Grip_colour       string  `db:"GRIP_COLOUR"`
	Nib_material      string  `db:"NIB_MATERIAL"`
	Nib_colour        string  `db:"NIB_COLOUR"`
	Nib_size          string  `db:"NIB_SIZE"`
	Purchase_price    float64 `db:"PURCHASE_PRICE"`
	Replacement_price float64 `db:"REPLACEMENT_PRICE"`
	Writing_type      string  `db:"WRITING_TYPE"`
	Cap_material      string  `db:"CAP_MATERIAL"`
	Cap_colour        string  `db:"CAP_COLOUR"`
	Image_1           string  `db:"IMAGE_1"`
	Image_2           string  `db:"IMAGE_2"`
	Comments          string  `db:"comments"`
}

type Polychrom struct {
	Item_id       int    `db:"item_id"`
	Key_Value     string `db:"key_value"`
	Colour_number string `db:"COLOUR_NUMBER"`
	Colour_name   string `db:"COLOUR_NAME"`
	Length        int    `db:"LENGTH"`
	Colour_group  string `db:"COLOUR_GROUP"`
	Location      string `db:"LOCATION"`
	RGB_red       int    `db:"RGB_RED"`
	RGB_green     int    `db:"RGB_GREEN"`
	RGB_blue      int    `db:"RGB_BLUE"`
	Web_name      string `db:"web_name"`
	Image_1       string `db:"IMAGE_1"`
	Comments      string `db:"comments"`
}

type Score struct {
	Item_id     int    `db:"item_id"`
	Key_Value   string `db:"key_value"`
	Publisher   string `db:"PUBLISHER"`
	Composer    string `db:"COMPOSER"`
	Work_name   string `db:"WORK_NAME"`
	Work_type   string `db:"WORK_TYPE"`
	Score_size  string `db:"SCORE_SIZE"`
	Score_pages int    `db:"SCORE_PAGES"`
	Comments    string `db:"comments"`
}

type Interface interface {
	GetOSMapList(ctx context.Context, mapRange string) ([]OSMap, error)
	GetOSMapItem(ctx context.Context, mapRange string, item_id int) (OSMap, error)
	GetInkList(ctx context.Context) ([]Ink, error)
	GetInkItem(ctx context.Context, item_id int) (Ink, error)
	GetPenList(ctx context.Context) ([]Pen, error)
	GetPenItem(ctx context.Context, item_id int) (Pen, error)
	GetScoreList(ctx context.Context) ([]Score, error)
	GetScoreItem(ctx context.Context, item_id int) (Score, error)
	GetPolychromList(ctx context.Context) ([]Polychrom, error)
	GetPolychromItem(ctx context.Context, item_id int) (Polychrom, error)
}
