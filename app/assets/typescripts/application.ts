/// <reference path="jquery.d.ts"/>
 let project = {
     name: "Gentry Rolofson",
     lang: "Golang ",
     version: 1.7
 }
 jQuery(function() {
     var myArray = project.lang.split("");
     jQuery(".navbar-brand").text(project.name);
     for (var i = 0; i < project.lang["length"]; i++) {
         jQuery(".lang").append("<span>" + myArray[i] + "</span>");
     }
     console.log((project.lang + " " + project.version));
     console.log(("Developed by: " + project.name));
     console.log(myArray[0]);
 });