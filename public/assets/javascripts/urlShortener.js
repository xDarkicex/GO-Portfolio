document.addEventListener("DOMContentLoaded", function(e){
  new Clipboard(".clipBoard")
})

document.addEventListener("submit", function(e){
  var xhr = new XMLHttpRequest();
  xhr.onload = function() {
     if (xhr.status != 200) {
       return
     }
     document.getElementsByClassName("results")[0].value = xhr.responseText
 }
 var original = document.getElementById("original").value
 document.getElementsByClassName("original")[0].innerHTML = original
 xhr.open("POST", "http://localhost:8080/");
 xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
 xhr.send("original="+original)
e.preventDefault()

})
