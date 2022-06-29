# wow-cgo

[![](https://img.shields.io/github/v/tag/thechampagne/wow-cgo?label=version)](https://github.com/thechampagne/wow-cgo/releases/latest) [![](https://img.shields.io/github/license/thechampagne/wow-cgo)](https://github.com/thechampagne/wow-cgo/blob/main/LICENSE)
### Installation & Setup

#### 1. Clone the repository
```
git clone https://github.com/thechampagne/wow-cgo.git
```
#### 2. Navigate to the root
```
cd wow-cgo
```
#### 3. Build the project
```
make
```
#### 4. Run test
```
make test
```

### Example

```c
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
```

### References
 - [wow](https://github.com/gernest/wow)

### License

This repo is released under the [MIT License](https://github.com/thechampagne/wow-cgo/blob/main/LICENSE).
