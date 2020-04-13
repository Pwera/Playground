var ul = document.querySelector("ul");

document.getElementById("add-btn").addEventListener("click", function(e) {
  e.preventDefault();

  var li = document.createElement("li"),
    pFirst = document.createElement("p"),
    pSecond = document.createElement("p"),
    iFirst = document.createElement("i"),
    iSecond = document.createElement("i"),
    input = document.createElement("input");
  addInput = document.getElementById("add-input");

  if (addInput.value === "") {
    return;
  }

  iFirst.className = "fa fa-pencil-square-o";
  iSecond.className = "fa fa-times";

  input.className = "edit note";
  input.setAttribute("type", "text");

  pFirst.textContent = addInput.value;

  pSecond.appendChild(iFirst);
  pSecond.appendChild(iSecond);
  li.appendChild(pFirst);
  li.appendChild(pSecond);
  li.appendChild(input);

  document.getElementById("list").appendChild(li);
  addInput.value = "";
});

ul.addEventListener("click", function(e) {
  if (e.target.classList[1] === "fa-pencil-square-o") {
    console.log("s");
    var parentPar = e.target.parentNode;
    parentPar.style.display = "none";

    var note = parentPar.previousElementSibling;
    var input = parentPar.nextElementSibling;

    input.style.display = "block";
    input.value = note.textContent;

    input.addEventListener("keypress", function(e) {
      if (e.keyCode === 13) {
        if (input.value !== "") {
          note.textContent = input.value;
          parentPar.style.display = "block";
          input.style.display = "none";
        } else {
          var li = input.parentNode;
          li.parentNode.removeChild(li);
        }
      }
    });
  } else if (e.target.classList[1] === "fa-times") {
    var list = e.target.parentNode.parentNode;
    list.parentNode.removeChild(list);
  }
});

var hideItem = document.getElementById("hide");
hideItem.addEventListener("click", function(e) {
  var label = document.querySelector("label");

  if (hideItem.checked) {
    label.textContent = "Unhide notes";
    ul.style.display = "none";
  } else {
    label.textContent = "Hide notes";
    ul.style.display = "block";
  }
});

var searchInput = document.querySelector("#search-note input");
searchInput.addEventListener("keyup", function(e) {
  var searchChar = e.target.value.toUpperCase();
  var notes = ul.getElementsByTagName("li");

  Array.from(notes).forEach(function(note) {
    var parText = note.firstElementChild.textContent.toUpperCase();
    if (parText.indexOf(searchChar) !== -1) {
      note.style.display = "clock";
    } else {
      note.style.display = "none";
    }
  });
});