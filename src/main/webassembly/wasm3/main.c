#include <stdio.h>
#include <emscripten.h>

// Declare reusable JS function
EM_JS(void, jsFunction, (int n)){
  console.log("Call from EM_JS" + n);
}

int main() {
  printf("WASM ready\n");

  // Call JS function eval
  emscripten_run_script("console.log('Hello from C')");
  // Call JS function async eval
  emscripten_async_run_script("console.log('Hello from C')", 2000);
  int intJsVal = emscripten_run_script_int("getInt()");
  char* stringJsVal = emscripten_run_script_string("getString()");
  printf("%d", intJsVal);

  //Execute  EM_JS
  jsFunction(123);
  return 42;
}

int getNum(){
  return 999;
}