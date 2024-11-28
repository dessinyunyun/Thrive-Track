package tools

import (
	"math"

	"github.com/gin-gonic/gin"
)

const (
	defaultLimit = 20
	maxLimit     = 1000
)

type Pagination struct {
	Limit        int `json:"limit" form:"limit"`
	Offset       int `json:"-"`
	Page         int `json:"page" form:"page" binding:"min=0"`
	NextPage     int `json:"next_page"`
	PreviousPage int `json:"previous_page"`
	Count        int `json:"count"`
	TotalPage    int `json:"total_page"`
}

// Paginate validates pagination requests
func Paginate(c *gin.Context) (*Pagination, error) {
	p := new(Pagination)
	if err := c.ShouldBindQuery(p); err != nil {
		return nil, err
	}
	if p.Page == 0 {
		p.Page = 1
	}
	if p.Limit < 1 {
		p.Limit = defaultLimit
	}
	if p.Limit > 1000 {
		p.Limit = maxLimit
	}
	p.Offset = (p.Page - 1) * p.Limit
	return p, nil
}

func Paging(p *Pagination) *Pagination {
	p.TotalPage = int(math.Ceil(float64(p.Count) / float64(p.Limit)))

	if p.Page > 1 {
		p.PreviousPage = p.Page - 1
	} else {
		p.PreviousPage = p.Page
	}

	if p.Page == p.TotalPage {
		p.NextPage = p.Page
	} else {
		p.NextPage = p.Page + 1
	}
	return p
}
