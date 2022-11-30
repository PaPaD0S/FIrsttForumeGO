// handlers
package hanldlers

import (
	"github/PaPa_D0S/FIrsttForumeGO/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "Nurmain.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "About.html")
}
