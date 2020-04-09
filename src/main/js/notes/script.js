// mutiple element
var icons = document.getElementsByClassName("fa");

// for(var i=0;i<icons.length; i++){
//     console.log(icons[i]);
// }


var iconsArr = Array.from(icons);
console.log(iconsArr);
console.log(icons);

iconsArr.forEach(function(icon, index, arr ) {
    console.log(icon, index, arr);
});

//single element
var el = document.querySelector("ul li:nth-child(4)")


// getElementsByTagName('tag') return a collection of all aelements in
// the document with the specified tag name

lis = document.getElementsByTagName('li');


// NodeList has forEach without transforming to Array

var aa = document.querySelectorAll("ul li")

aa[0].style.cssText = "font-size: 25px";

// className get and sets the value of class attribute of the specified element

var h2 = document.querySelector("header h2");

h2.addEventListener("click", function(e){
    console.log("clicked", e);
});