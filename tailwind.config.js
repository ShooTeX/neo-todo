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
    boxShadow: {
      default: "8px 5px 0px 0px #000000",
      sm: "2px 3px 0px 0px #000000",
      lg: "10px 7px 0px 0px #000000",
    }
  },
  plugins: [],
}

