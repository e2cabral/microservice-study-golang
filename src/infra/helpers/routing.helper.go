package helpers

type Router struct {
	Group string
}

func NewRouter(group string) *Router {
	return &Router{Group: group}
}
