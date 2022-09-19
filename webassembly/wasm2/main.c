#include <string.h>

void consoleLog(int n);
void strLog(char* msg, int n);

int main() { 
  return 42;
}

void myfunc(int n){
  return consoleLog(23);
}

void greet(){
  char* msg = "Hello from C";
  strLog(msg, strlen(msg));
}