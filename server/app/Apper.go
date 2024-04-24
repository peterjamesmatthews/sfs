package app

type App struct {
	db Databaser
}

func New(db Databaser) *App {
	return &App{}
}
