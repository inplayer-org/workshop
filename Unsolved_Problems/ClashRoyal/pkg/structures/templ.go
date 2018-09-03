package structures

import "html/template"


//Tmpl - Global template
var Tmpl = template.Must(template.ParseGlob("../tmpl/*.html"))

