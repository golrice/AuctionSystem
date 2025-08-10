package kernal

type Pagination struct {
	Page uint `form:"page" json:"page" default:"1" validate:"min=1"`
	Size uint `form:"size" json:"size" default:"10" validate:"min=1,max=100"`
}

func (p *Pagination) Offset() uint {
	return (p.Page - 1) * p.Size
}

func (p *Pagination) Limit() uint {
	return p.Size
}

func NewDefaultPagination() Pagination {
	return Pagination{
		Page: 1,
		Size: 10,
	}
}
