# declutter

Declutter is a tool to organize your files in correct folder structure. 
Remove all the clutter from your desktop or any folders from your terminal!

Declutter looks into all the files(excluding hidden files) provided in a given directory. Then based on extension type and configuration provided 
on config.yaml file, moves those files into correct folders.

## Configure output folder 
Change output foldername and add different type based on your personal need in config.yaml file. 

For example:
```
output:
  trash: 
    - txt
  photos:
    - png
    - jpg

```

## Usage

Run go build 

```bash
$ go build 
```

Run the declutter executable and pass the dir path you want to organize.

```bash
$ ./declutter ~/Desktop
```
