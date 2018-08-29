package structures

import "html/template"

var Tmpl = template.Must(template.ParseGlob("../tmpl/tableranking.html"))

