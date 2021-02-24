var items = document.querySelectorAll(".item");
var points = document.querySelectorAll(".point");
var areas = document.querySelectorAll("area");
var live = document.querySelector('.live')
var goIndex = function (Index) {
  clearactive();
  items[Index].className = "item active";
  points[Index].className = "point active";
};
var clearactive = function () {
  for (let i = 0; i < items.length; i++) {
    items[i].className = "item";
    points[i].className = "point";
  }
};
for (let i = 0; i < points.length; i++) {
  points[i].addEventListener("click", function () {
    var Index = this.getAttribute("data-index");
    goIndex(Index);
  });
}
$("areas").hover(
  function (params) {
    $("areas").css("background-color", "yellow");
  },
  function () {
    $("areas").css("background-color", "red");
  }
);
$(live).scrollTop(20);
