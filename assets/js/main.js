// Bootstrap JS — modular imports
// Mirrors assets/scss/main.scss: uncomment components as needed.
// Each import pulls only that component + its dependencies via esbuild tree-shaking.

// import Alert from 'bootstrap/js/dist/alert';
// import Button from 'bootstrap/js/dist/button';
// import Carousel from 'bootstrap/js/dist/carousel';
// import Collapse from 'bootstrap/js/dist/collapse';
// import Dropdown from 'bootstrap/js/dist/dropdown';
import Modal from 'bootstrap/js/dist/modal';
// import Offcanvas from 'bootstrap/js/dist/offcanvas';
// import Popover from 'bootstrap/js/dist/popover';
// import ScrollSpy from 'bootstrap/js/dist/scrollspy';
// import Tab from 'bootstrap/js/dist/tab';
// import Toast from 'bootstrap/js/dist/toast';
// import Tooltip from 'bootstrap/js/dist/tooltip';

// Copy-to-clipboard for elements with data-copy-target
document.addEventListener('click', (e) => {
  const btn = e.target.closest('[data-copy-target]');
  if (!btn) return;
  const target = document.getElementById(btn.dataset.copyTarget);
  if (!target) return;
  navigator.clipboard.writeText(target.textContent.trim()).then(() => {
    const icon = btn.querySelector('i');
    if (icon) {
      icon.className = 'bi bi-check2';
      setTimeout(() => { icon.className = 'bi bi-clipboard'; }, 1500);
    }
  });
});

// Email signup (layouts/partials/subscribe-modal.html). Submit the form to Kit
// in the background and show Kit's answer in the modal. Kit's endpoint allows
// cross-origin posts and answers JSON with a status of "success",
// "quarantined" (Kit wants a bot check on its own page), or "failed" with error
// messages. If the request itself fails, fall back to a plain form post.
const subscribeForm = document.querySelector('[data-subscribe-form]');
if (subscribeForm) {
  const modal = document.getElementById('subscribeModal');
  const status = modal.querySelector('[data-subscribe-status]');
  const note = modal.querySelector('[data-subscribe-note]');
  const button = subscribeForm.querySelector('button[type="submit"]');

  const say = (text, kind) => {
    status.textContent = text;
    status.className = `subscribe-status mt-3 mb-0${kind ? ` text-${kind}` : ''}`;
  };

  subscribeForm.addEventListener('submit', async (e) => {
    e.preventDefault();
    const body = new FormData(subscribeForm);
    body.append('referrer', document.referrer);
    body.append('host', window.location.href);
    button.disabled = true;
    say('Subscribing…');

    let data;
    try {
      const res = await fetch(subscribeForm.action, {
        method: 'POST',
        body,
        headers: { Accept: 'application/json' },
      });
      if (!res.ok) throw new Error(res.statusText);
      data = await res.json();
    } catch {
      subscribeForm.submit();
      return;
    }

    button.disabled = false;
    if (data.status === 'success') {
      subscribeForm.hidden = true;
      note.hidden = true;
      say('Thanks! Now check your email to confirm your subscription.', 'success');
    } else if (data.status === 'quarantined' && data.url) {
      say('');
      const link = document.createElement('a');
      link.href = data.url;
      link.target = '_blank';
      link.rel = 'noopener';
      link.textContent = 'confirm on Kit’s site';
      status.append('Kit needs to check that you are not a bot. Please ', link, ' to finish subscribing.');
    } else {
      const messages = data.errors?.messages ?? [];
      say(messages[0] || 'Something went wrong. Please try again.', 'danger');
    }
  });

  // Reset the form when the modal closes, so it starts fresh next time.
  modal.addEventListener('hidden.bs.modal', () => {
    subscribeForm.reset();
    subscribeForm.hidden = false;
    note.hidden = false;
    button.disabled = false;
    say('');
  });
}
