package request

type List struct {
	Search  string `json:"search"`
	Sort    string `json:"sort"`
	Order   string `json:"order"`
	PerPage int    `json:"per_page"`
	Page    int    `json:"page"`
}
