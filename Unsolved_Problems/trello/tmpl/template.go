package template

import "html/template"

var Tmpl = template.Must(template.ParseGlob("../tmpl/*.html"))
