package handlers

import (
	"net/http"

	"github.com/agungdhewe/dwlog"
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/defaulthandlers"
)

func (hdr *Handler) Login(w http.ResponseWriter, r *http.Request) {
	ws := hdr.Webservice
	hdr.Webservice.Session.RenewToken(r.Context())
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)

	var showloginpage = false
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			if ws.ShowServerError {
				defaulthandlers.ErrorPageHandler(500, err.Error(), pv, w, r)
			} else {
				dwlog.Error(err.Error())
				defaulthandlers.ErrorPageHandler(500, "internal server error", pv, w, r)
			}
			return
		}

		email := r.Form.Get("email")
		password := r.Form.Get("password")

		var authenticated bool
		authenticated = false
		if email == "agung" && password == "rahasia" {
			authenticated = true
		}

		if authenticated {
			// login berhasil, redirect ke halaman yang di refer sebelumnya
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {
			showloginpage = true
		}

	} else {
		showloginpage = true
	}

	if showloginpage {
		pv.Form = appsmodel.NewForm(r.PostForm)
		pv.PageName = "login"
		defaulthandlers.SimplePageHandler(pv, w, r)
	}

}