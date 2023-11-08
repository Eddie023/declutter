# declutter

![logo](docs/resources/logo.png)

Declutter is a tool to organize your files in correct folder structure.
Remove all the clutter from your specified directory from your terminal!

Declutter looks into all the files(excluding hidden files) provided in a given directory. Then based on extension type moves those files into relevant folders.

## Usage:

```
Usage: declutter [options...] <filepath>

Options:
        -v Show verbose logs. (WIP)
        -r Show what would the output look like without moving files. (WIP)
```

## Build:

```bash
$ make build
```

Run the declutter executable and pass the dir path you want to organize.

```bash
$ ./declutter ~/Desktop
```
