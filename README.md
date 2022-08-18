# digrep

![Github Actions](https://github.com/autopp/digrep/actions/workflows/push-main.yml/badge.svg)

digrep is filter command by `.dockerignore`.

Like `grep`, `digrep` takes input from standard input and and outputs non filtered line. Unlike `grep`, `digrep` filter by `.dockerignore` (matched lines are OMITTED).

## Install

Download from [releases page](https://github.com/autopp/digrep/releases).

## Usage

E.g.
```
$ ls | digrep
```

## License

[Apache License 2.0](LICENSE)
