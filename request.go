package opentaobao

type Requester interface {
	GetMethod() string
	GetApiParams() Parameter
}
