package cli

type App interface {
	Run() error
}

func New() App {
	return newAppv()
}

type appv struct {
	cwd string
}

func newAppv() *appv {
	return nil
}

func (a *appv) Run() error {
	return nil
}
