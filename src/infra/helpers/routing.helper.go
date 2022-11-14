package helpers

type Router struct {
	Prefix string
}

func NewRouter(prefix string) *Router {
	return &Router{Prefix: prefix}
}
