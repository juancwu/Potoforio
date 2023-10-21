/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['./public/views/**/*.html'],
    theme: {
        extend: {
            fontFamily: {
                inter: ['Inter', 'sans-serif'],
            },
            screens: {
                '3xl': '2048px',
                '4xl': '2560px',
                '5xl': '4096px',
            },
        },
    },
    plugins: [],
};
