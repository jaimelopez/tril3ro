#ifndef __MODULE_DARWIN_H__
#define __MODULE_DARWIN_H__

#include <mach-o/dyld_images.h>
#include <mach-o/loader.h>
#include <sys/stat.h>
#include <mem_darwin.h>

#define PATH_MAX 2048

typedef struct Module
{
  mach_vm_address_t addr;
  off_t size;
  const char *module;
} Module;


struct Module *get_modules(pid_t pid, uint32_t *count);
#endif
