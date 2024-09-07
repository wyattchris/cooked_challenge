import type { Config } from "tailwindcss";
import { nextui } from "@nextui-org/react";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
    "./node_modules/@nextui-org/theme/dist/components/button.js",
    './node_modules/@nextui-org/theme/dist/components/(button|snippet|code|input).js'
  ],
  theme: {
    extend: {
      colors: {
        primary: "#D9D9D9",
        primaryHover: "#C9C9C9",
        secondary: "#393939",
        secondaryHover: "#292929",
        tertiary: "#000000"
      }
    },
  },
  darkMode: "class",
  plugins: [nextui()],
};
export default config;
