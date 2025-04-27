document.addEventListener("DOMContentLoaded", function () {
  setTimeout(() => {
    Prism.highlightAll();
    console.log("WORKING...");
  }, 1000);
});

window.addEventListener("hashchange", function () {
  Prism.highlightAll(); // Re-run Prism highlighting when URL changes
});
