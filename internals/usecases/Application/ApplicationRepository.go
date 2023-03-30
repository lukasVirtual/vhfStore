package application

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

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

// Gethers all information of all available Applications
// for the gui to display. Currently fetches JSON file.
func (a *ApplicationRepository) FetchCatalog() {
	log.Debug("ApplicationRepository::FetchCatalog")

	http.Get("http://127.0.0.1:3333/api/prepare")
	resp, err := http.Get("http://127.0.0.1:8000/infos.json")
	if err != nil {
		//TODO Error Handling
		log.Fatal(err)
	}

	data, err := os.ReadFile("catalog.json")
	if err != nil {
		//TODO Error Handling
		log.Fatal(err)
	}


	log.Println(data, resp)
}

// Downloads a single Application through choosing the right download strategy
func (a *ApplicationRepository) Download(name string) {
	log.Debug("ApplicationRepository::Download")
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:3333/%v.json", name))
	if err != nil {
		//TODO Error Handling
		log.Fatal(err)
	}

	log.Println(resp.Body)
}

// Installs the Application with multiple
func (a *ApplicationRepository) Install(name string) {
	log.Debug("ApplicationRepository::Install")
}

func (a *ApplicationRepository) Uninstall(name string) {
	log.Debug("ApplicationRepository::Uinstall")
}

func (a *ApplicationRepository) Update(name string) {
	log.Debug("ApplicationRepository::Update")
}
