//main javascript
$ project = {
    name: "Gentry Rolofson",
    lang: "Golang",
    version: 1.7
}

jQuery(fn ():
    jQuery(".navbar-brand").text(project.name)
    jQuery(".lang").prepend(project.lang+" ")
    say (project.lang +" "+ project.version)
    say ("Developed by: " + project.name)
end)

