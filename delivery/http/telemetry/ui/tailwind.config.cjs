/* eslint-disable-next-line */
const colors = require('tailwindcss/colors')

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './index.html',
    './src/**/*.{js,ts,jsx,tsx}'
  ],
  theme: {
    extend: {
      colors: {
        safe: colors.green,
        warning: colors.amber,
        danger: colors.red
      }
    }
  },
  plugins: []
}
