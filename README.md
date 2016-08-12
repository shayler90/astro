# astro

[ ![Codeship Status for dnathanturner/astro](https://codeship.com/projects/f640d060-4300-0134-8a1c-1e76a446a475/status?branch=master)](https://codeship.com/projects/168312) [![Coverage Status](https://coveralls.io/repos/bitbucket/dnathanturner/astro/badge.svg?branch=master)](https://coveralls.io/bitbucket/dnathanturner/astro?branch=master)

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
