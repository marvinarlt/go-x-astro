import { defineConfig } from 'astro/config';

// https://astro.build/config
/** @type {import('astor/config').AstroUserConfig} */
export default defineConfig({
  outDir: './../dist',
  build: {
    format: 'directory',
  },
});
