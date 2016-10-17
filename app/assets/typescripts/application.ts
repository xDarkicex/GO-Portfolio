/// <reference path="jquery.d.ts"/>
let project = {
     name: "Gentry Rolofson",
     lang: [ "Golang Developer", "Ruby on Rails Developer", "NodeJS Developer", "Angular Developer", "Bash Scripter"],
     devicon: [ '<span class="devicons devicons-go"></span>', 
     '<span class="devicons devicons-ruby_on_rails"></span>',
     '<span class="devicons devicons-nodejs"></span>',
     '<span class="devicons devicons-angular"</span>',
     '<span class="devicons devicons-terminal"</span>' ],
     version: 1.7
}

let di = project.devicon
let v = project.lang
let k = 0

jQuery(function() {
    jQuery(".navbar-brand").text(project.name)
    langText()
    setInterval(langText, 3000);
})

 function langText(){
    $(".lang").html(v[k])
    $("span.devicons").html(di[k])
    k++
    if (k > v.length-1){
        k = 0
    }
 }