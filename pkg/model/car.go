package model

type Car struct {
	Id          int    `json:"id"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
	Image_url   string `json:"image_url"`
	Year        int    `json:"year"`
	Price       string `json:"price"`
	Color       string `json:"color"`
	Description string `json:"description"`
}
