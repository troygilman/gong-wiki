/** @type {import('tailwindcss').Config} */
module.exports = {
    content: {
        files: ["./**/*.templ"],
    },
    theme: {
        extend: {},
    },
    plugins: [require("@tailwindcss/typography"), require("daisyui")],
};
