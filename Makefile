all: clean build compile run

clean:
	rm -rf ./build ./reactjs/build
	echo "[✔️] Clean complete!"

build:
	cd ./reactjs && npm install
	cd ./reactjs && npm run build
	echo "[✔️] Build complete!"

compile:
	mkdir -p ./build
	go build -o build
	echo "[✔️] Compile complete!"


run:
	./build/go-react-app
	echo "[✔️] App is running!"
