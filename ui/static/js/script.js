// --- Theme Toggler ---
const themeToggle = document.getElementById("theme-toggle");
const body = document.body;

// Function to set theme
const setTheme = (theme) => {
  body.setAttribute("data-theme", theme);
};

// Event listener for the button
themeToggle.addEventListener("click", () => {
  const currentTheme = body.getAttribute("data-theme");
  if (currentTheme === "dark") {
    setTheme("light");
  } else {
    setTheme("dark");
  }
});

// --- Smooth Scrolling for Navigation Links ---
document.querySelectorAll('a[href^="#"]').forEach((anchor) => {
  anchor.addEventListener("click", function (e) {
    e.preventDefault();
    document.querySelector(this.getAttribute("href")).scrollIntoView({
      behavior: "smooth",
    });
  });
});
