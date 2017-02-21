/// <reference path="../jquery.d.ts"/>
$(function(){
$(document).scroll(function(){
    var $window = $(window)
    var $boxes = $('.box')
    if ($boxes.length > 0) {
        $boxes.each(function(i, e){
            let $box = $(e)
            let $container = $box.parent().parent()
            let offset = Math.max(0, $window.scrollTop() - $container.offset().top + 75)
            offset = Math.min($container.height() + $container.offset().top - $box.height() * 2, offset)
            let currentOffset = parseFloat($box[0].style.top)
            if(offset !== currentOffset) $box.css({top:offset})
        })
    }
})
})
