# declutter

![logo](docs/resources/logo.png)

Declutter is a tool to organize your files in correct folder structure. 
Remove all the clutter from your specified directory all this from your terminal!

Declutter looks into all the files(excluding hidden files) provided in a given directory. Then based on extension type moves those files into relevant folders.

## Usage

Run go build 

```bash
$ go build 
```

Run the declutter executable and pass the dir path you want to organize.
If no path is provided then the current directory is assumed as the folder that needs cleaning.

```bash
$ ./declutter ~/Desktop
```
