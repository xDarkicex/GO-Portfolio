/// <reference path="jquery.d.ts"/>

function shrinker() {
    
 }

 var xhr = new XMLHttpRequest();

 xhr.onload = function() {
    if (xhr.status === 200) {
        var data = JSON.parse(xhr.responseText);
        var input = document.getElementById('searchInput');
        var list = document.getElementById('json-datalist');
        var i: number;
        for (i = 0; i < data.length; i++) {
            var option = document.createElement('option');
            list.appendChild(option)
            option.value = data[i].Title;
            input.placeholder = "Search..."; 
        }
    }
}
xhr.open("GET", "/api/posts/search");
xhr.send()

$(function() {
    var $input = $('#searching');
    var $searchButton = $('#btn-search')
    $searchButton.on('click', function(){
        if($input.hasClass("shrinking")) $input.removeClass('shrinking')
        else $searchButton.parent().submit()

    })
})
