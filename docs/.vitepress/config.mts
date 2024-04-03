import { defineConfig } from 'vitepress';

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: 'Toodaloo',
  titleTemplate: ':title | Toodaloo',
  description: 'Say goodbye to your todos',
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Getting started', link: '/purpose' },
      { text: 'Reference', link: '/commands' },
      { text: 'Contributing', link: '/contributing' },
    ],

    editLink: {
      pattern: 'https://github.com/mrsimonemms/toodaloo/edit/main/docs/:path',
    },

    footer: {
      message: 'Released under Apache 2.0 Licence',
      copyright:
        'A project by <a href="https://simonemms.com" target="_blank">Simon Emms</a>',
    },

    sidebar: [
      {
        text: 'Getting started',
        items: [
          { text: 'Purpose of the project', link: '/purpose' },
          { text: 'Installation', link: '/install' },
        ],
      },
      {
        text: 'Reference',
        items: [
          { text: 'Commands', link: '/commands' },
          { text: 'Pre-commit hook', link: '/pre-commit' },
        ],
      },
      {
        text: 'Contributing',
        link: '/contributing',
      },
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/mrsimonemms/toodaloo' },
      { icon: 'twitter', link: 'https://twitter.com/theshroppiebeek' },
      { icon: 'linkedin', link: 'https://www.linkedin.com/in/simonemms' },
    ],
  },
});
