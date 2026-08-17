const API = {
  async get(url) {
    const res = await fetch(url);
    if (!res.ok) throw new Error((await res.json()).error || '请求失败');
    return res.json();
  },
  async post(url, data) {
    const res = await fetch(url, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data)
    });
    if (!res.ok) throw new Error((await res.json()).error || '请求失败');
    return res.json();
  },
  async put(url, data) {
    const res = await fetch(url, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data)
    });
    if (!res.ok) throw new Error((await res.json()).error || '请求失败');
    return res.json();
  },
  async delete(url) {
    const res = await fetch(url, { method: 'DELETE' });
    if (!res.ok) throw new Error((await res.json()).error || '请求失败');
    return res.json();
  },
  async uploadImage(file) {
    const fd = new FormData();
    fd.append('image', file);
    const res = await fetch('/api/upload', {
      method: 'POST',
      body: fd
    });
    if (!res.ok) throw new Error((await res.json()).error || '上传失败');
    return res.json();
  }
};

function showToast(message, type = 'success') {
  const colors = {
    success: 'bg-green-500',
    error: 'bg-red-500',
    info: 'bg-blue-500'
  };
  const toast = document.createElement('div');
  toast.className = `fixed top-20 right-4 ${colors[type]} text-white px-5 py-3 rounded-lg shadow-lg z-[9999] transition-all duration-300 translate-x-full`;
  toast.textContent = message;
  document.body.appendChild(toast);
  requestAnimationFrame(() => toast.classList.remove('translate-x-full'));
  setTimeout(() => {
    toast.classList.add('translate-x-full');
    setTimeout(() => toast.remove(), 300);
  }, 2500);
}

function openModal(contentHtml) {
  const backdrop = document.createElement('div');
  backdrop.className = 'modal-backdrop';
  backdrop.innerHTML = contentHtml;
  backdrop.addEventListener('click', (e) => {
    if (e.target === backdrop) backdrop.remove();
  });
  document.body.appendChild(backdrop);
  backdrop.querySelectorAll('[data-close-modal]').forEach(el => {
    el.addEventListener('click', () => backdrop.remove());
  });
  return backdrop;
}

function closeModal(backdrop) {
  if (backdrop) backdrop.remove();
  else document.querySelectorAll('.modal-backdrop').forEach(el => el.remove());
}

function formatDate(dateStr) {
  if (!dateStr) return '-';
  try {
    const d = new Date(dateStr);
    return d.toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' });
  } catch { return dateStr; }
}

function escapeHtml(str) {
  const div = document.createElement('div');
  div.textContent = str || '';
  return div.innerHTML;
}

function setActiveNav(page) {
  document.querySelectorAll('.nav-links a').forEach(a => {
    a.classList.toggle('active', a.dataset.page === page);
  });
}
