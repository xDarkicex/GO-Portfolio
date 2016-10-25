/// <reference path="jquery.d.ts"/>
var project = {
    name: "Gentry Rolofson",
    lang: [['Golang Developer', '<span class="devicons devicons-go"></span>'],
        ['Ruby on Rails Developer', '<span class="devicons devicons-ruby_on_rails"></span>'],
        ['NodeJS Developer', '<span class="devicons devicons-nodejs_small"></span>'],
        ['Angular Developer', '<span class="devicons devicons-angular"</span>'],
        ['Bash Scripter', '<span class="devicons devicons-terminal"</span>']],
    version: 1.7
};
var v = project.lang;
var i = 0;
jQuery(function () {
    jQuery(".navbar-brand").text(project.name);
    langText();
    setInterval(langText, 3000);
});
function langText() {
    $(".lang").html(v[i][0]);
    $("span.devicons").html(v[i][1]);
    i++;
    if (i > v.length - 1) {
        i = 0;
    }
}
function shrinker() {
    $("#searching").toggleClass("shrinking");
}
