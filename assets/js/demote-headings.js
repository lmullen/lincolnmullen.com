// Add the h3 class to all h1 tags, and so forth
// Effectively visually demotes all headings by two steps
const demoteBy = 2;
for (let i = 1; i <= 4; i++) {
  const current = "h" + i;
  const desired = "h" + (i + demoteBy);
  const headings = document.querySelectorAll(current);
  for (let heading of headings) {
    heading.classList.add(desired);
  }
}
