#include <stdio.h>
#include <unistd.h>
#include "wow.h"

int main()
{
  wow_t* wow;
  spinner_t* spin;
  spinner_t s;
  char* arr[1];
  spin = spin_get(DOTS);
  wow = wow_init(spin, "Such Spins");
  wow_start(wow);
  sleep(2);
  wow_text(wow, "Very emojis");
  wow_spinner(wow, spin_get(HEARTS));
  sleep(2);
  arr[0] = "👍"; 
  s.frames = arr;
  s.frames_length = 1;
  wow_persist_with(wow, &s,"Wow!");
  wow_clean(wow);
  spinner_clean(spin);
}
