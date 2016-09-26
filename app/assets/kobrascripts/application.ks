//main javascript
$ project = {
    name: "Gentry Rolofson",
    lang: "Golang ",
    version: 1.7
}

jQuery(fn ():
    $ myArray = project.lang.split("")
    jQuery(".navbar-brand").text(project.name)
    for ($ i = 0; i < project.lang["length"]; i++):
        jQuery(".lang").append("<span>"+myArray[i]+"</span>")
    end
    say (project.lang +" "+ project.version)
    say ("Developed by: " + project.name)
    
    say myArray[0]
end)
