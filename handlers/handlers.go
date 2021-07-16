package handlers

import (
	"github.com/CloudyKit/jet/v6"
	"log"
	"net/http"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./templates/html"),
	jet.InDevelopmentMode(), // remove in production
)

func Home(w http.ResponseWriter, r *http.Request) {
	err:= renderPage(w, "home.jet", nil)
	if err != nil {
		log.Println(err)
	}

}

func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	views, err:= views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}
	views.Execute(w, data, nil)
	return nil


}
