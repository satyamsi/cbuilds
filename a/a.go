// +build linux 

package a

/*
#include "a.h"
*/
import "C"

func Somefunc() (ret int) {

        x := C.somefunc() 
  
        ret = int(x) + 1

	return ret 
}
