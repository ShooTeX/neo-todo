const defaultTheme = require('tailwindcss/defaultTheme')

/** @type {import('tailwindcss').Config} */
export default {
  content: ["./web/views/**/*.html"],
  theme: {
    extend: {
      fontFamily: {
        'sans': ['"Space Mono"', ...defaultTheme.fontFamily.mono],
      },
    },
  },
  plugins: [],
}

