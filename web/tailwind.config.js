/** @type {import('tailwindcss').Config} */
export default {
  content: [
    
    "./src/**/*.jsx",
  ],
  theme: {
   
    extend: {
      colors : {
        'taupe': '#D1BFA7',
        'beige-taupe': '#B8A394',
        'sage-green': '#6B705C',
        'steel-blue': '#2A4D69'
      },
    },
  },
  plugins: [],
}

// "./src/**/*.{js,ts,jsx,tsx}",