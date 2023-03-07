package application

type ApplicationInterface interface {
	Download(string)
	Install(string)
	Uninstall(string)
	Update(string)
}

type ApplicationRepository struct {
}

func New() *ApplicationRepository {
	return &ApplicationRepository{}
}
