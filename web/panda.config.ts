import {
  defineConfig,
  defineSemanticTokens,
  defineTokens,
} from "@pandacss/dev";

import { button } from "src/theme/components/Button/button.recipe";
import { input } from "src/theme/components/Input/input.recipe";
import { link } from "src/theme/components/Link/link.recipe";
import { menu } from "src/theme/components/Menu/menu.recipe";

export default defineConfig({
  preflight: true,
  strictTokens: true,
  include: ["./src/**/*.tsx"],
  jsxFramework: "react",
  exclude: [],

  // NOTE: The theme references some CSS variables defined in global.css, this
  // is in order to provide some level of customisability for hosts who want to
  // override CSS with custom rules. Panda is primarily just there to wire it
  // all together and define the semantic tokens.
  //
  // NOTE: There may be some parts of global.css that reference tokens generated
  // by Panda, this is best avoided but it's some leftovers from the early days.
  theme: {
    recipes: {
      input: input,
      button: button,
      link: link,
      menu: menu,
    },
    extend: {
      semanticTokens: defineSemanticTokens({
        colors: {
          accent: {
            50: { value: "var(--accent-colour-flat-fill-50)" },
            100: { value: "var(--accent-colour-flat-fill-100)" },
            200: { value: "var(--accent-colour-flat-fill-200)" },
            300: { value: "var(--accent-colour-flat-fill-300)" },
            400: { value: "var(--accent-colour-flat-fill-400)" },
            500: { value: "var(--accent-colour-flat-fill-500)" },
            600: { value: "var(--accent-colour-flat-fill-600)" },
            700: { value: "var(--accent-colour-flat-fill-700)" },
            800: { value: "var(--accent-colour-flat-fill-800)" },
            900: { value: "var(--accent-colour-flat-fill-900)" },
            text: {
              50: { value: "var(--accent-colour-flat-text-50)" },
              100: { value: "var(--accent-colour-flat-text-100)" },
              200: { value: "var(--accent-colour-flat-text-200)" },
              300: { value: "var(--accent-colour-flat-text-300)" },
              400: { value: "var(--accent-colour-flat-text-400)" },
              500: { value: "var(--accent-colour-flat-text-500)" },
              600: { value: "var(--accent-colour-flat-text-600)" },
              700: { value: "var(--accent-colour-flat-text-700)" },
              800: { value: "var(--accent-colour-flat-text-800)" },
              900: { value: "var(--accent-colour-flat-text-900)" },
            },
            dark: {
              50: { value: "var(--accent-colour-dark-fill-50)" },
              100: { value: "var(--accent-colour-dark-fill-100)" },
              200: { value: "var(--accent-colour-dark-fill-200)" },
              300: { value: "var(--accent-colour-dark-fill-300)" },
              400: { value: "var(--accent-colour-dark-fill-400)" },
              500: { value: "var(--accent-colour-dark-fill-500)" },
              600: { value: "var(--accent-colour-dark-fill-600)" },
              700: { value: "var(--accent-colour-dark-fill-700)" },
              800: { value: "var(--accent-colour-dark-fill-800)" },
              900: { value: "var(--accent-colour-dark-fill-900)" },
              text: {
                50: { value: "var(--accent-colour-dark-text-50)" },
                100: { value: "var(--accent-colour-dark-text-100)" },
                200: { value: "var(--accent-colour-dark-text-200)" },
                300: { value: "var(--accent-colour-dark-text-300)" },
                400: { value: "var(--accent-colour-dark-text-400)" },
                500: { value: "var(--accent-colour-dark-text-500)" },
                600: { value: "var(--accent-colour-dark-text-600)" },
                700: { value: "var(--accent-colour-dark-text-700)" },
                800: { value: "var(--accent-colour-dark-text-800)" },
                900: { value: "var(--accent-colour-dark-text-900)" },
              },
            },
          },
          bg: {
            canvas: { value: "{colors.gray.100}" },
            default: {
              value: { base: "{colors.white}", _dark: "{colors.gray.200}" },
            },
            subtle: {
              value: { base: "{colors.gray.200}", _dark: "{colors.gray.300}" },
            },
            muted: {
              value: { base: "{colors.gray.300}", _dark: "{colors.gray.400}" },
            },
            emphasized: {
              value: { base: "{colors.gray.400}", _dark: "{colors.gray.500}" },
            },
            disabled: {
              value: { base: "{colors.gray.300}", _dark: "{colors.gray.400}" },
            },
          },
          fg: {
            default: { value: "{colors.gray.900}" },
            muted: { value: "{colors.gray.800}" },
            subtle: { value: "{colors.gray.700}" },
            disabled: { value: "{colors.gray.500}" },
          },
          border: {
            default: { value: "{colors.gray.400}" },
            muted: { value: "{colors.gray.500}" },
            subtle: { value: "{colors.gray.300}" },
            disabled: { value: "{colors.gray.400}" },

            outline: { value: "{colors.blackAlpha.50}" },
            accent: { value: "{colors.accent.default}" },
          },
        },
      }),
      tokens: {
        fontSizes: {
          sm: { value: "1rem" },
          md: { value: "1.125rem" },
          lg: { value: "1.2rem" },
          xl: { value: "1.44rem" },
          "2xl": { value: "1.728rem" },
          "3xl": { value: "2.074rem" },
          "4xl": { value: "2.488rem" },
          heading: {
            1: { value: "var(--font-size-h1)" },
            2: { value: "var(--font-size-h2)" },
            3: { value: "var(--font-size-h3)" },
            4: { value: "var(--font-size-h4)" },
            5: { value: "var(--font-size-h5)" },
            6: { value: "var(--font-size-h6)" },
            variable: {
              1: { value: "var(--font-size-h1-variable)" },
              2: { value: "var(--font-size-h2-variable)" },
              3: { value: "var(--font-size-h3-variable)" },
              4: { value: "var(--font-size-h4-variable)" },
              5: { value: "var(--font-size-h5-variable)" },
              6: { value: "var(--font-size-h6-variable)" },
            },
          },
        },
        colors: defineTokens.colors({
          whiteAlpha: {
            50: { value: "rgba(255, 255, 255, 0.04)" },
            100: { value: "rgba(255, 255, 255, 0.06)" },
            200: { value: "rgba(255, 255, 255, 0.08)" },
            300: { value: "rgba(255, 255, 255, 0.16)" },
            400: { value: "rgba(255, 255, 255, 0.24)" },
            500: { value: "rgba(255, 255, 255, 0.36)" },
            600: { value: "rgba(255, 255, 255, 0.48)" },
            700: { value: "rgba(255, 255, 255, 0.64)" },
            800: { value: "rgba(255, 255, 255, 0.80)" },
            900: { value: "rgba(255, 255, 255, 0.92)" },
          },
          blackAlpha: {
            50: { value: "rgba(0, 0, 0, 0.04)" },
            100: { value: "rgba(0, 0, 0, 0.06)" },
            200: { value: "rgba(0, 0, 0, 0.08)" },
            300: { value: "rgba(0, 0, 0, 0.16)" },
            400: { value: "rgba(0, 0, 0, 0.24)" },
            500: { value: "rgba(0, 0, 0, 0.36)" },
            600: { value: "rgba(0, 0, 0, 0.48)" },
            700: { value: "rgba(0, 0, 0, 0.64)" },
            800: { value: "rgba(0, 0, 0, 0.80)" },
            900: { value: "rgba(0, 0, 0, 0.92)" },
          },
        }),
      },
    },
  },

  outdir: "styled-system",
});
