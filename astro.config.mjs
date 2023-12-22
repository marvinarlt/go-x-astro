import { defineConfig } from 'astro/config';

// https://astro.build/config
/** @type {import('astor/config').AstroUserConfig} */
export default defineConfig({
  srcDir: './client',
  build: {
    format: 'file',
  },
});
