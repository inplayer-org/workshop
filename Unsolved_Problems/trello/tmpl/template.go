package tmpl

import "html/template"

var Tmpl = template.Must(template.ParseGlob("../tmpl/src/*.html"))
