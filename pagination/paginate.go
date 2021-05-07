package pagination

import (
	"fmt"
	"gorm.io/gorm"
	"os"
)

// Result ..
type Result struct {
	CurrentPage  int         `json:"current_page"`
	Data         interface{} `json:"data"`
	FirstPageURL string      `json:"first_page_url"`
	From         int         `json:"from"`
	LastPage     int         `json:"last_page"`
	LastPageURL  string      `json:"last_page_url"`
	NextPageURL  string      `json:"next_page_url"`
	Path         string      `json:"path"`
	PerPage      int         `json:"per_page"`
	PrevPageURL  string      `json:"prev_page_url"`
	To           int         `json:"to"`
	Total        int64       `json:"total"`
}

// Paginator ..
type Paginator interface {
	Paginate(db *gorm.DB) (Result, *gorm.DB)
}

// Config ..
//
// If you don't provide app url it will fetch the APP_URL from environment
type Config struct {
	Page    int
	Sort    string
	PerPage int
	AppURL  string
	Path    string
}

// Paginate ..
func (c *Config) Paginate(db *gorm.DB, any interface{}) (Result, *gorm.DB) {
	var r Result
	var count int64

	offset := (c.Page - 1) * c.PerPage
	lastIndex := offset * c.Page
	d := db.Offset(offset).Limit(c.PerPage)

	if c.Sort != "" {
		d.Order(c.Sort)
	}

	d.Find(any)

	db.Model(any).Count(&count)

	r.CurrentPage = c.Page
	r.NextPageURL = c.GetPageURL(c.Page + 1)
	r.FirstPageURL = c.GetPageURL(1)
	r.PrevPageURL = c.PreviousPageURL()
	r.PerPage = c.PerPage
	r.Path = c.Path
	r.To = lastIndex
	r.From = lastIndex - offset
	r.Total = count
	r.Data = any
	r.LastPageURL = c.GetPageURL(r.GetLastPage())
	r.LastPage = r.GetLastPage()

	return r, d
}

// GetLastPage ..
func (r *Result) GetLastPage() int {
	return int(r.Total) / r.PerPage
}

func (c *Config) GetPageURL(page int) string {
	return fmt.Sprintf("%s%s?page=%d&per_page=%d", c.GetAppURL(), c.Path, page, c.PerPage)
}

// PreviousPageURL ..
func (c *Config) PreviousPageURL() string {
	pageNumber := 1

	if c.Page > 1 {
		pageNumber = c.Page - 1
	}

	return c.GetPageURL(pageNumber)
}

// GetAppURL ..
func (c *Config) GetAppURL() string {
	if c.AppURL == "" {
		return os.Getenv("APP_URL")
	}

	return c.AppURL
}
