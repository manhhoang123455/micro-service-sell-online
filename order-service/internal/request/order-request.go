package request

type OrderRequest struct {
	Items []OrderItem `json:"items" required:"true"`
}

type OrderItem struct {
	ProductID uint `json:"product_id" required:"true"`
	Quantity  int  `json:"quantity" required:"true"`
}
