/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['./src/**/*.{html,js,svelte,ts}'],
    plugins: [
        require('@tailwindcss/typography'),
        require("daisyui")
    ],
    daisyui: {
        logs: false,
        darkTheme: "business",
        themes: ["corporate", "business"],
    },
}
