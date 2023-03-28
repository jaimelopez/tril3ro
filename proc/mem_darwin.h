#ifndef __MEM_DARWIN_H__
#define __MEM_DARWIN_H__

#include <stdio.h>
#include <mach-o/dyld_images.h>
#include <mach/vm_map.h>

uintptr_t readProcessMemory(int pid, mach_vm_address_t addr, mach_msg_type_number_t *size);

bool writeProcessMemory(int pid, mach_vm_address_t addr, vm_offset_t data, mach_msg_type_number_t size);
#endif