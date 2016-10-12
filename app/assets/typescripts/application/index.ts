/// <reference path="../jquery.d.ts"/>
$(function(){
    console.log(" I worked!!")
    $("#arrows").click(function() {
    $('html,body').animate({
        scrollTop: $("#scrolled").offset().top},
        'slow');
});
})

