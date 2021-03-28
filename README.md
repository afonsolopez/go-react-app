# Go React App

This project is a boilerplate for building React apps for Linux, MacOS and Windows systems, through webviews. This project is an adaptation of Vic Shóstak's repository [example-go-react-macos-app-1](https://github.com/koddr/example-go-react-macos-app-1).

## Installation

Clone this project into your machine.

```bash
git clone https://github.com/afonsolopez/go-react-app.git
```
Install if you're running OpenBSD, FreeBSD or Ubuntu "webkit2gtk" as the [webview docs](https://github.com/webview/webview#notes) says. As I am running Xubuntu 20.04 on my machine this following package worked fine. 

```bash
sudo apt-get install libwebkit2gtk-4.0-dev
```
Install all needed Go modules packages.
```go
go get -u github.com/gobuffalo/packr
go get -u github.com/webview/webview
```
## Usage

```bash
# Open the cloned project folder
cd go-react-app

# Clean the project, build the React app, compiles the code and runs the app.
make

# Just cleans older builds
make clean

# Builds the React app
make build

# Compiles the app
make compile

# Runs the compiled app
make run
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)