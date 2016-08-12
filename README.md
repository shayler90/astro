# astro

[ ![Codeship Status for dnt17/astro](https://codeship.com/projects/4af7d1a0-4303-0134-bd17-76dd7bf7b13e/status?branch=master)](https://codeship.com/projects/168314) [![Coverage Status](https://coveralls.io/repos/github/dnt17/astro/badge.svg?branch=master)](https://coveralls.io/github/dnt17/astro?branch=master)

**astro** is a command line application designed to automate the deployment process of applications. It will parse a config file (.astro.toml), which will identify:

  - the application
  - the source repo
  - the version
  - if astro should pull the master branch and push a newly created branch upstream
    - if set to false, astro will pull the branch that it is currently on
  - optional extra dependencies
  - the build steps
  - the deployment environment
  - any post build steps

### Configuration structure

Examples:

  - [basic](examples/basic_example.toml)
  - [medium](examples/medium_example.toml)
