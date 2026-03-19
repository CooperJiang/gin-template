/** @type {import('tailwindcss').Config} */
export default {
  darkMode: 'class',
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {
      fontFamily: {
        sans: [
          'Space Grotesk',
          'Inter',
          'ui-sans-serif',
          'system-ui',
          '-apple-system',
          'BlinkMacSystemFont',
          'sans-serif',
        ],
      },
      boxShadow: {
        'nb': '4px 4px 0px var(--nb-border)',
        'nb-sm': '2px 2px 0px var(--nb-border)',
        'nb-lg': '6px 6px 0px var(--nb-border)',
        'nb-hover': '6px 6px 0px var(--nb-border)',
        'nb-active': '1px 1px 0px var(--nb-border)',
      },
      borderWidth: {
        'nb': '2.5px',
      },
      colors: {
        nb: {
          primary: 'var(--nb-primary)',
          'primary-hover': 'var(--nb-primary-hover)',
          bg: 'var(--nb-bg)',
          'bg-soft': 'var(--nb-bg-soft)',
          surface: 'var(--nb-surface)',
          'surface-alt': 'var(--nb-surface-alt)',
          text: 'var(--nb-text)',
          'text-secondary': 'var(--nb-text-secondary)',
          'text-muted': 'var(--nb-text-muted)',
          border: 'var(--nb-border)',
          pink: 'var(--nb-accent-pink)',
          blue: 'var(--nb-accent-blue)',
          orange: 'var(--nb-accent-orange)',
          purple: 'var(--nb-accent-purple)',
          cyan: 'var(--nb-accent-cyan)',
          red: 'var(--nb-accent-red)',
          green: 'var(--nb-accent-green)',
        },
      },
    },
  },
  plugins: [require('tailwindcss-animate')],
}
