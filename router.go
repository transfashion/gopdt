package main

import (
	"github.com/fgtago/fgweb"
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/defaulthandlers"
	"github.com/fgtago/gopdt/handlers"
	"github.com/go-chi/chi/v5"
)

func Router(mux *chi.Mux) error {

	fgweb.Get(mux, "/favicon.ico", defaulthandlers.FaviconHandler)
	fgweb.Get(mux, "/asset/*", defaulthandlers.AssetHandler)
	fgweb.Get(mux, "/template/*", defaulthandlers.TemplateHandler)

	hnd := handlers.New(appsmodel.GetWebservice())

	// Home
	fgweb.Get(mux, "/", hnd.Home)
	fgweb.Get(mux, "/manifest.json", hnd.Manifest)

	// Pricing
	fgweb.Get(mux, "/module/pricing", hnd.Pricing)
	fgweb.Post(mux, "/module/pricing", hnd.PricingPost)

	// login activity
	fgweb.Get(mux, "/user/account", hnd.Account)
	fgweb.Get(mux, "/user/logout", hnd.Logout)
	fgweb.Get(mux, "/user/login", hnd.Login)
	fgweb.Post(mux, "/user/login", hnd.LoginPost)

	// Standard Module
	fgweb.Get(mux, "/module/{modulename}", hnd.Module)

	// fgweb.Get(mux, "/{any:.*}", func(w http.ResponseWriter, r *http.Request) {
	// 	param := chi.URLParam(r, "any")
	// 	fmt.Print(param)
	// 	http.NotFound(w, r)
	// })

	return nil
}
