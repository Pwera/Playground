#include <stdio.h>
#include <time.h>
#include <stdlib.h>
#include <emscripten.h>

#define NUM_CIRCLES 100

struct Circle {
  int x;
  int y;
  int r;
  int cr;
  int cg;
  int cb;  
};

struct CircleAnimationData {
  int x;
  int y;
  int r;
  int xv; // axis velocity
  int yv; // axis velocity
  int xd; // axis direction- bool
  int yd; // axis direction- bool
  
};

struct Circle circles[NUM_CIRCLES];
struct CircleAnimationData animationData[NUM_CIRCLES];

struct Circle* getCircles(int canvasWidth, int canvasHeight){
    //Update circle data

  for (int i=0; i<NUM_CIRCLES; i++){

    //Collision right 
    if((animationData[i].x + animationData[i].r) >= canvasWidth) {
      animationData[i].xd=0;
    }     

    //Collision left 
    if((animationData[i].x - animationData[i].r) <= 0){
      animationData[i].xd=1;
    } 
      

    //Collision top
    if((animationData[i].y - animationData[i].r) <= 0) {
      animationData[i].yd=1;
    }
      
    //Collision bottom
    if((animationData[i].y + animationData[i].r) >= canvasHeight){
      animationData[i].yd=0;
    }

    //Move circle in specified direction
    if(animationData[i].xd ==1){
      animationData[i].x +=animationData[i].xv;
    }else{
      animationData[i].x -=animationData[i].xv;
    }

    if(animationData[i].yd ==1){
      animationData[i].y +=animationData[i].yv;
    }else{
      animationData[i].y -=animationData[i].yv;
    }

    //Update matching circle
    circles[i].x=animationData[i].x;
    circles[i].y=animationData[i].y;

  }
  return circles;
}
int getRand(int max){
  return (rand() % max);
}

int main() {
  srand(time(NULL));
  for (int i=0; i< NUM_CIRCLES; i++)  {
    int radius = getRand(50);

    int x = getRand(1000) + radius;
    int y = getRand(1000) + radius;

    animationData[i].x = x;
    animationData[i].y = y;
    animationData[i].r = radius;
    animationData[i].xv=getRand(10);
    animationData[i].yv=getRand(10);
    animationData[i].xd=1;
    animationData[i].yd=1;

    circles[i].x = x;
    circles[i].y = y;
    circles[i].r = radius;
    circles[i].cr = getRand(255);
    circles[i].cg = getRand(255);
    circles[i].cb = getRand(255);
  }

  // Start JS rendering
  EM_ASM({render($0, $1);}, NUM_CIRCLES * 6, 6);

}

