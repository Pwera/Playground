yarn init
yarn add express
yarn start

emcc main.c  -s WASM=1 -o public/demo.js
emcc main.c  -s WASM=1 -o public/demo.html
//produce only a wasm
emcc main.c  -s WASM=1 -s SIDE_MODULE=1 -o public/demo.wasm
//optymalization
emcc main.c  -s WASM=1 -O2 -o public/demo.js
// export functions
emcc main.c  -s WASM=1 -s EXPORTED_FUNCTIONS="['_getNum','_main']" -o public/demo.js
// ccall comes from preamble.js
ccall('getNum')

# Emrun support
emcc main.c  -s WASM=1 --emrun -o public/index.html
emrun --port 8888 --no_browser public/index.html

emcc main.c  -s WASM=1 -s EXPORTED_FUNCTIONS="['_getCircles','_main']" -o public/demo.js && yarn start