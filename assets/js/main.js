tailwind.config = {
    theme: {
        extend: {
            fontFamily: {
                display: ['Outfit', 'sans-serif'],
                mono: ['JetBrains Mono', 'monospace'],
            },
            colors: {
                nex: {
                    900: '#020a1a',
                    800: '#041030',
                    700: '#0a1a3f',
                    600: '#0f2452',
                    500: '#1a3a7a',
                    400: '#2563eb',
                    300: '#3b82f6',
                    200: '#60a5fa',
                    100: '#93c5fd',
                    50: '#dbeafe',
                    glow: '#4f8bff',
                    cyan: '#06d6e0',
                    violet: '#7c3aed',
                }
            }
        }
    }
}



const particleContainer = document.getElementById('particles');
function createParticle() {
    const p = document.createElement('div');
    p.className = 'particle';
    const size = Math.random() * 3 + 1;
    const colors = ['rgba(59,130,246,0.5)', 'rgba(6,214,224,0.4)', 'rgba(124,58,237,0.4)'];
    p.style.cssText = `
                width: ${size}px; height: ${size}px;
                left: ${Math.random() * 100}%;
                bottom: -10px;
                background: ${colors[Math.floor(Math.random() * colors.length)]};
                animation-duration: ${Math.random() * 15 + 10}s;
                animation-delay: ${Math.random() * 5}s;
            `;
    particleContainer.appendChild(p);
    setTimeout(() => p.remove(), 25000);
}
setInterval(createParticle, 800);


const starContainer = document.getElementById('stars');
for (let i = 0; i < 60; i++) {
    const s = document.createElement('div');
    s.className = 'star';
    const size = Math.random() * 1.5 + 0.5;
    s.style.cssText = `
                width: ${size}px; height: ${size}px;
                left: ${Math.random() * 100}%;
                top: ${Math.random() * 100}%;
                animation-duration: ${Math.random() * 4 + 2}s;
                animation-delay: ${Math.random() * 3}s;
            `;
    starContainer.appendChild(s);
}


document.querySelectorAll('.glow-card').forEach(card => {
    card.addEventListener('mousemove', e => {
        const rect = card.getBoundingClientRect();
        card.style.setProperty('--mouse-x', `${e.clientX - rect.left}px`);
        card.style.setProperty('--mouse-y', `${e.clientY - rect.top}px`);
    });
});


const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
        if (entry.isIntersecting) {
            entry.target.style.animationPlayState = 'running';
        }
    });
}, { threshold: 0.1 });

document.querySelectorAll('.stagger-in').forEach(el => {
    el.style.animationPlayState = 'paused';
    observer.observe(el);
});
