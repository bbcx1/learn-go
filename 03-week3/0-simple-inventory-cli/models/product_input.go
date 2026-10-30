package models

type ProductInput struct {
	Name  string
	Price uint
}

func (p *ProductInput) Validate() bool {
	return p.Name != "" && p.Price >= 1000
}
