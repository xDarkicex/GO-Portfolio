/// <reference path="../jquery.d.ts"/>
$(function(){
    $(".card-block > p").addClass("truncate")
    $("#arrows").click(function() {
    $('html,body').animate({
        scrollTop: $("#scrolled").offset().top},
        'slow');
});
})

