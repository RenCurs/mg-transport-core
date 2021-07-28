// Copyright (c) 2019 RetailDriver LLC
// Use of this source code is governed by a MIT
/*
Package core provides different functions like error-reporting, logging, localization, etc. in order to make it easier to create transports.
Usage:
	package main

	import (
	    "os"
	    "fmt"
	    "html/template"

	    "github.com/gin-gonic/gin"
	    "github.com/retailcrm/mg-transport-core/core"
	)

	func main() {
	    app := core.New()
	    app.Config = core.NewConfig("config.yml")
	    app.DefaultError = "unknown_error"
	    app.TranslationsPath = "./translations"

	    app.ConfigureRouter(func(engine *gin.Engine) {
	        engine.Static("/static", "./static")
	        engine.HTMLRender = app.CreateRenderer(
	            func(renderer *core.Renderer) {
	                // insert templates here. Example:
	                r.Push("home", "templates/layout.html", "templates/home.html")
	            },
	            template.FuncMap{},
	        )
	    })

	    if err := app.Prepare().Run(); err != nil {
	        fmt.Printf("Fatal error: %s", err.Error())
	        os.Exit(1)
	    }
	}

Resource embedding

packr can be used to provide resource embedding, see:
	https://github.com/gobuffalo/packr/tree/master/v2
In order to use packr you must follow instruction, and provide boxes with templates, translations and assets to library.
You can find instruction here:
	https://github.com/gobuffalo/packr/tree/master/v2#library-installation
Example of usage:
	package main

	import (
	    "os"
	    "fmt"
	    "html/template"
		"embed"
		"net/http"

	    "github.com/gin-gonic/gin"
	    "github.com/retailcrm/mg-transport-core/core"
	)

	//go:embed static
	var Static embed.FS

	//go:embed translations
	var TranslationsFS embed.FS
	var TranslationsDir string

	//go:embed templates
	var TemplatesFS embed.FS
	var TemplatesDir string

	func main() {
	    app := core.New()
	    app.Config = core.NewConfig("config.yml")
	    app.DefaultError = "unknown_error"
	    app.TranslationsDir = TranslationsDir
	    app.TranslationsFS = TranslationsFS

	    app.ConfigureRouter(func(engine *gin.Engine) {
	        engine.StaticFS("/assets", http.FS(Static))
	        engine.HTMLRender = app.CreateRendererFS(
	            TemplatesFS,
	            TemplatesDir,
	            func(renderer *core.Renderer) {
	                // insert templates here. Example:
	                r.Push("home", "layout.html", "home.html")
	            },
	            template.FuncMap{},
	        )
	    })

	    if err := app.Prepare().Run(); err != nil {
	        fmt.Printf("Fatal error: %s", err.Error())
	        os.Exit(1)
	    }
	}

Migration generator

This library contains helper tool for transports. You can install it via go:
	$ go get -u github.com/retailcrm/mg-transport-core/cmd/transport-core-tool
Currently, it only can generate new migrations for your transport.
*/
package core
