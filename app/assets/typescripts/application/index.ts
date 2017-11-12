/// <reference path="../jquery.d.ts"/>
$(function(){
    $("#arrows").click(function() {
    $('html,body').animate({
        scrollTop: $("#scrolled").offset().top},
        'slow');
});
})

let project = {
     lang: [ ['Golang Developer', '<span class="devicons devicon-go-line"></span>'],
     ['Ruby on Rails Developer', '<span class="devicons devicon-rails-plain-wordmark"></span>'], 
    ['NodeJS Developer', '<span class="devicons devicon-nodejs-plain"></span>'], 
    ['Angular Developer', '<span class="devicons devicon-angularjs-plain-wordmark"></span>'], 
    ['React Developer', '<span class="devicons devicon-react-original"></span>'],
    ['Amazon Cloud Solution Architect', '<span class="devicons devicon-amazonwebservices-plain-wordmark"</span>' ]],
     version: 2.0
}
let v = project.lang
let i = 0

jQuery(function() {
    langText()
    setInterval(langText, 3000);
})

 function langText(){
    $(".lang").html(v[i][0])
    $("span.devicons").html(v[i][1])
    i++
    if (i > v.length-1){
        i = 0
    }
 }