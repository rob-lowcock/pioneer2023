/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'webscale': '#6298BC',
        'webscale-lighter': '#70A1C2',
        'bggray': '#F3F3F3',
        'brdgray': '#DDDEDF',
        'black': '#15171a',
        'slate': '#7C8B9A',
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}

