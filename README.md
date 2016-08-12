# astro

**astro** is a command line application designed to automate the deployment process of Think Research applications. It will parse a config file (.astro.toml), which will identify:

  - the application
  - the source repo
  - the version
  - if astro should pull the master branch and push a newly created branch upstream
    - if set to false, astro will pull the branch that it is currently on
  - optional extra dependencies
  - the build steps
  - the deployment environment
  - any post build steps

Configurations can be overwritten using [flaegs](https://github.com/containous/flaeg).

### Configuration structure

Examples:

  - [basic](examples/basic_example.toml)
