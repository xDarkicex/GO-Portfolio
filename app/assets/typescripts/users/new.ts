/// <reference path="../jquery.d.ts"/>
function isValid() {
    var password = $("#Password").val()
    var confirmPassword = $("#Password2").val()
    if (password.length >= 8){
        checkPasswordMatch()
    }   else {
             $("#divCheckPasswordMatch").html("Password must be at least 8 characters!")

    }
}
function checkPasswordMatch() {
    var password = $("#Password").val()
    var confirmPassword = $("#Password2").val()

    if (password != confirmPassword) {
        $("#divCheckPasswordMatch").html("Passwords do not match!")
        $("#registration > div > div.form-group").removeClass("has-success")
        $("#registration > div > div.form-group").addClass("has-danger")
        $("#Password, #Password2").removeClass("form-control-success")
        $("#Password, #Password2").addClass("form-control-danger")
        $("#isValid").attr('disabled', 1);
        
    } else {
        $("#divCheckPasswordMatch").html("Passwords match.")
        $("#registration > div > div.form-group").removeClass("has-danger")
        $("#registration > div > div.form-group").addClass("has-success")
        $("#Password, #Password2").removeClass("form-control-danger")
        $("#Password, #Password2").addClass("form-control-success")
         $("#isValid").removeAttr('disabled');
    }
}

$(document).ready(function () {
    // checkPasswordExists()
    $("#Password2, #Password").keyup(function(){
        isValid()
    })
})
