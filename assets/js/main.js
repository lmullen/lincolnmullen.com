// Bootstrap JS — modular imports
// Mirrors assets/scss/main.scss: uncomment components as needed.
// Each import pulls only that component + its dependencies via esbuild tree-shaking.

// import Alert from 'bootstrap/js/dist/alert';
// import Button from 'bootstrap/js/dist/button';
// import Carousel from 'bootstrap/js/dist/carousel';
// import Collapse from 'bootstrap/js/dist/collapse';
// import Dropdown from 'bootstrap/js/dist/dropdown';
// import Modal from 'bootstrap/js/dist/modal';
// import Offcanvas from 'bootstrap/js/dist/offcanvas';
// import Popover from 'bootstrap/js/dist/popover';
// import ScrollSpy from 'bootstrap/js/dist/scrollspy';
// import Tab from 'bootstrap/js/dist/tab';
// import Toast from 'bootstrap/js/dist/toast';
import Tooltip from 'bootstrap/js/dist/tooltip';

// Initialize tooltips on all elements with data-bs-toggle="tooltip"
document.querySelectorAll('[data-bs-toggle="tooltip"]').forEach(el => {
  new Tooltip(el);
});
