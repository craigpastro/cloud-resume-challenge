var counterContainer = document.querySelector(".counter");
var visitCount = localStorage.getItem("page_view");

if (!visitCount) {
    visitCount = 0
}

visitCount = Number(visitCount) + 1;
localStorage.setItem("page_view", visitCount)
counterContainer.innerHTML = visitCount;
