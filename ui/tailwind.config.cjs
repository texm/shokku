module.exports = {
  content: ["./src/**/*.{html,js,svelte}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    darkTheme: "business",
    themes: ["corporate", "business", {
      lightish: {
        "primary": "#034280",
        "secondary": "#d5ba46",
        "accent": "#ea6747",
        "neutral": "#001523",
        "base-100": "#FDFDFD",
        "info": "#4A82E3",
        "success": "#24BC70",
        "warning": "#F7C655",
        "error": "#E52461",
      },
      /*
        Tandoori #b75c50
        Moonlit Forest #36716a
        Continental Waters #94c0c4
       */
    }],
    logs: false,
  },
}
