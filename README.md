# CLS (C Language Support)
Current verion v1.00.00

This is a little cli application for make c/c++ experience more comfortable.
This application i made mostly for myself and for few of my friends. So don't expect much

## Installation
I didn't use any librares so you just need to install golang
Debian/Ubuntu
```bash
sudo apt install golang
```
Fedora
```bash
sudo dnf install golang
```
Arch
```bash
sudo pacman -S go
```

And then you could
```bash
git clone https://github.com/Fufik-pivufik/C-language-support
cd C-languange-support
./install
```
And then check 
```bash
cls version
```

# Usage
When you did all previous steps you can use *cls*

```bash
cls help # command for view all options
```




## New
Create new project
```bash
cls new <project_name> # this commad creates a simple hello world app with basic structure
```
### flags:
 -c makes new C project instead of C++

## Biuld & Run
Build your project or build and run it with one command.
You should run it somewhere in your project dir: src, include, etc
```bash
cls build
```
Run
```bash
cls run 
```
### flags
-h hides all unnecessary information(files, executable file, etc)
-r build project without flags from field flags in cls.json

## Config
Config file cls.json main file for your project, all necessary information contains here
Show your current configuration
```bash
cls config show
```
Change name of your project. Doesn't change directory name? But build your project with new name
```bash
cls config name <new_name>
```

## Flag
You can manually add flags for compilation and remove them
```bash
cls flag add <flag>
```
```bash
cls flag remove <flag>
```
All flags contain in cls.json


## Get
You can download library from github. 
```bash
cls get <URL>
```
After that you should install it manually in
<project_path>/extend/<library>
Because so many libraries have their own way to installation


## Test
Before creating test you should check the name of your main file and change it to your current file with main function
If main file in config matches with your actual main file you can create test
### Create
```bash
cls test create
```
It'll make directory test/ in your project's home folder

### Run
```bash
cls test run
```



