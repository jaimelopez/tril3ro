#ifndef __MODULE_DARWIN_H__
#define __MODULE_DARWIN_H__

#include <stdio.h>
#include <stdlib.h>
#include <mach-o/dyld_images.h>
#include <mach-o/loader.h>
#include <mach/vm_map.h>

#define PATH_MAX 2048

typedef struct Module { 
  mach_vm_address_t addr;
  uint32_t size;
  const char * module;
} Module;

struct Module * getModules(pid_t pid, uint32_t* count);
unsigned char* readProcessMemory(int pid, mach_vm_address_t addr, mach_msg_type_number_t *size);
#endif