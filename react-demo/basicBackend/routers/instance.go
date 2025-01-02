package routers

type Gateway struct {
	id int64
}

func NewGateway(id int64) *Gateway {
	return &Gateway{
		id: id,
	}
}
