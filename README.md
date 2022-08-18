# digrep

![Github Actions](https://github.com/autopp/digrep/actions/workflows/push-main.yml/badge.svg)

digrep is filter command by `.dockerignore`.

Like `grep`, `digrep` takes input from standard input and and outputs non filtered line. Unlike `grep`, `digrep` filter by `.dockerignore` (matched lines are OMITTED).

## Install

Download from [releases page](https://github.com/autopp/digrep/releases).

## Usage

When no parameter is given, `digrep` reads `.dockerignore` from current directory

E.g.
```
$ digrep < files.txt
```

When parameter is given, `digrep` assumes it is a directory and reads `.dockerignore` from it.
```
$ digrep docker < files.txt
```

When `.dockerignore` is not found, `digrep` outputs all input lines.

## License

[Apache License 2.0](LICENSE)
