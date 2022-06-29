package main

/*
#include <stdlib.h>

typedef enum {
  UNKNOWN = 0,
  ARC,
  ARROW,
  ARROW2,
  ARROW3,
  BALLOON,
  BALLOON2,
  BOUNCE,
  BOUNCINGBALL,
  BOUNCINGBAR,
  BOXBOUNCE,
  BOXBOUNCE2,
  CHRISTMAS,
  CIRCLE,
  CIRCLEHALVES,
  CIRCLEQUARTERS,
  CLOCK,
  DOTS,
  DOTS10,
  DOTS11,
  DOTS12,
  DOTS2,
  DOTS3,
  DOTS4,
  DOTS5,
  DOTS6,
  DOTS7,
  DOTS8,
  DOTS9,
  DQPB,
  EARTH,
  FLIP,
  GRENADE,
  GROWHORIZONTAL,
  GROWVERTICAL,
  HAMBURGER,
  HEARTS,
  LAYER,
  LINE,
  LINE2,
  MONKEY,
  MOON,
  NOISE,
  PIPE,
  POINT,
  PONG,
  RUNNER,
  SHARK,
  SIMPLEDOTS,
  SIMPLEDOTSSCROLLING,
  SMILEY,
  SQUARECORNERS,
  SQUISH,
  STAR,
  STAR2,
  TOGGLE,
  TOGGLE10,
  TOGGLE11,
  TOGGLE12,
  TOGGLE13,
  TOGGLE2,
  TOGGLE3,
  TOGGLE4,
  TOGGLE5,
  TOGGLE6,
  TOGGLE7,
  TOGGLE8,
  TOGGLE9,
  TRIANGLE,
  WEATHER,
} name_t;

typedef struct {
  name_t name;
  int interval;
  int frames_length;
  char** frames;
} spinner_t;

typedef struct {
  int _id;
  int is_terminal;
} wow_t;
*/
import "C"
import (
	"os"
	"unsafe"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
)

var (
  wowMap = make(map[int]*wow.Wow)
  wowMapLen = 0
)

//export wow_init
func wow_init(s *C.spinner_t, text *C.char) *C.wow_t {
  self := (*C.wow_t) (C.malloc(C.size_t(unsafe.Sizeof(C.wow_t{}))))
  var frames []string
  slice := (*[1 << 30]*C.char)(unsafe.Pointer(s.frames))[:s.frames_length:s.frames_length]
  for _, v := range slice {
    frames = append(frames, C.GoString(v))
  }
  spin := spin.Spinner{
    Name: spin.Name(s.name),
    Interval: int(s.interval),
    Frames: frames,
  }
  w := wow.New(os.Stdout, spin, C.GoString(text))
  wowMapLen += 1
  wowMap[wowMapLen] = w
  if w.IsTerminal {
    self._id = C.int(wowMapLen)
    self.is_terminal = 1
    return self
  }
    self._id = C.int(wowMapLen)
    self.is_terminal = 0
    return self
}

//export wow_persist
func wow_persist(self *C.wow_t) {
  wowMap[int(self._id)].Persist()
}

//export wow_persist_with
func wow_persist_with(self *C.wow_t, s *C.spinner_t, text *C.char) {
  var frames []string
  slice := (*[1 << 30]*C.char)(unsafe.Pointer(s.frames))[:s.frames_length:s.frames_length]
  for _, v := range slice {
    frames = append(frames, C.GoString(v))
  }
  spin := spin.Spinner{
    Name: spin.Name(s.name),
    Interval: int(s.interval),
    Frames: frames,
  }
  wowMap[int(self._id)].PersistWith(spin, C.GoString(text))
}

//export wow_spinner
func wow_spinner(self *C.wow_t, s *C.spinner_t) {
  var frames []string
  slice := (*[1 << 30]*C.char)(unsafe.Pointer(s.frames))[:s.frames_length:s.frames_length]
  for _, v := range slice {
    frames = append(frames, C.GoString(v))
  }
  spin := spin.Spinner{
    Name: spin.Name(s.name),
    Interval: int(s.interval),
    Frames: frames,
  }
  w := wowMap[int(self._id)].Spinner(spin)
  wowMap[int(self._id)] = w
  if w.IsTerminal {
    self.is_terminal = 1
    return
  }
    self.is_terminal = 0
    return
}

//export wow_start
func wow_start(self *C.wow_t) {
  wowMap[int(self._id)].Start()
}

//export wow_stop
func wow_stop(self *C.wow_t) {
  wowMap[int(self._id)].Stop()
}

//export wow_text
func wow_text(self *C.wow_t, txt *C.char) {
  w := wowMap[int(self._id)].Text(C.GoString(txt))
  wowMap[int(self._id)] = w
  if w.IsTerminal {
    self.is_terminal = 1
    return
  }
    self.is_terminal = 0
    return
}

//export spin_get
func spin_get(name C.name_t) *C.spinner_t {
  self := (*C.spinner_t) (C.malloc(C.size_t(unsafe.Sizeof(C.spinner_t{}))))
  sn := spin.Get(spin.Name(name))
  array := C.malloc(C.size_t(len(sn.Frames)) * C.size_t(unsafe.Sizeof(uintptr(0))))
  slice := (*[1<<30 - 1]*C.char)(array)
  for i, v := range sn.Frames {
    slice[i] = C.CString(v)
  }
  self.name = C.name_t(sn.Name)
  self.interval = C.int(sn.Interval)
  self.frames = (**C.char) (array)
  self.frames_length = C.int(len(sn.Frames))
  return self
}

//export wow_clean
func wow_clean(self *C.wow_t) {
  delete(wowMap, int(self._id))
  if self != nil {
    C.free(unsafe.Pointer(self))
  }
}

//export spinner_clean
func spinner_clean(self *C.spinner_t) {
  if self != nil {
    slice := (*[1 << 30]*C.char)(unsafe.Pointer(self.frames))[:self.frames_length:self.frames_length]
    for _, v := range slice {
      if v != nil {
        C.free(unsafe.Pointer(v))
      }
    }
    C.free(unsafe.Pointer(self.frames))
    C.free(unsafe.Pointer(self))
  }
}

func main() {}