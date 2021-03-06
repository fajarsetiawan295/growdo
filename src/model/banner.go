package model

import (
	"fmt"
	"log"
	"os"
	"time"
)

// "id" bigserial PRIMARY KEY,
//     "images" varchar NOT NULL,
//     "status" BOOLEAN NOT NULL DEFAULT (false),

type Banner struct {
	Id         uint64    `json:"id"`
	Images     string    `json:"images" validate:"required"`
	Status     bool      `json:"status" validate:"required"`
	Url        string    `json:"url"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type funcBanner struct {
	Data  string
	Model *Banner
	Cari  *FilterCari
}

func NewBanner(data string, M *Banner, C *FilterCari) *funcBanner {
	return &funcBanner{
		Data:  data,
		Model: M,
		Cari:  C,
	}
}

var (
	columBanner     = "images, status, url"
	tabelBanner     = "banner"
	selectImgBanner = fmt.Sprintf("CONCAT('%s', images), status, url", os.Getenv("BASE_URL"))
)

func (r *funcBanner) Create() string {
	// data := "Banner () VALUES ($1, $2, $3, $4, $5, $6)"

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ('%s', %v, '%s') RETURNING ID",
		tabelBanner,
		columBanner,
		r.Model.Images,
		r.Model.Status,
		r.Model.Url,
	)
	log.Println(query)
	return query
}

// Userfindemail = "email=$1"
func (r *funcBanner) Detail() string {

	var hasil string
	var data = fmt.Sprintf("SELECT id, %s, %s FROM %s", selectImgBanner, DateGlobal(), tabelBanner)

	switch r.Data {
	case "status":
		hasil = fmt.Sprintf("%s WHERE status=%t", data, r.Cari.StatusBanner)
	case "id":
		hasil = fmt.Sprintf("%s WHERE id=%d", data, r.Cari.Id)
	case "all":
		hasil = fmt.Sprintf("%s ORDER BY id DESC", data)
	}

	log.Println(hasil)
	return hasil
}
