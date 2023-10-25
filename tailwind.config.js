/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./pkg/pages/**/*.go", "./public/views/**/*.html"],
    theme: {
        extend: {
            fontFamily: {
                inter: ["Inter", "sans-serif"],
            },
            screens: {
                "3xl": "2048px",
                "4xl": "2560px",
                "5xl": "4096px",
                // for mac users
                retina: {
                    raw: `@media
only screen and (-webkit-min-device-pixel-ratio: 2),
only screen and (min--moz-device-pixel-ratio: 2),
only screen and (-o-min-device-pixel-ratio: 2/1),
only screen and (min-device-pixel-ratio: 2),
only screen and (min-resolution: 192dpi),
only screen and (min-resolution: 2dppx)
`,
                },
                "retina-sm": {
                    raw: `@media
only screen and (-webkit-min-device-pixel-ratio: 2) and (min-width: 640px),
only screen and (min--moz-device-pixel-ratio: 2) and (min-width: 640px),
only screen and (-o-min-device-pixel-ratio: 2/1) and (min-width: 640px),
only screen and (min-device-pixel-ratio: 2) and (min-wdith: 640px),
only screen and (min-resolution: 192dpi) and (min-width: 640px),
only screen and (min-resolution: 2dppx) and (min-width: 640px)
`,
                },
                "retina-md": {
                    raw: `@media
only screen and (-webkit-min-device-pixel-ratio: 2) and (min-width: 1024px),
only screen and (min--moz-device-pixel-ratio: 2) and (min-width: 1024px),
only screen and (-o-min-device-pixel-ratio: 2/1) and (min-width: 1024px),
only screen and (min-device-pixel-ratio: 2) and (min-wdith: 1024px),
only screen and (min-resolution: 192dpi) and (min-width: 1024px),
only screen and (min-resolution: 2dppx) and (min-width: 1024px)
`,
                },
            },
        },
    },
    plugins: [],
};
