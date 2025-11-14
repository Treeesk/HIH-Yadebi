import forms from "@tailwindcss/forms"
import lineClamp from "@tailwindcss/line-clamp"
import aspectRatio from "@tailwindcss/aspect-ratio"

export default {
    content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
    theme: { extend: {} },
    plugins: [
        forms(),
        lineClamp(),
        aspectRatio(),
    ],
}