#ifndef __MEM_DARWIN_H__
#define __MEM_DARWIN_H__

#include <stdlib.h>
#include <mach-o/dyld_images.h>
#include <mach/vm_map.h>

uintptr_t read_process_memory(task_t task, mach_vm_address_t addr, mach_msg_type_number_t *size);

void read_process_memory_bytes(task_t task, mach_vm_address_t addr, void *buffer, mach_msg_type_number_t *size);

bool write_process_memory(task_t task, mach_vm_address_t addr, vm_offset_t data, mach_msg_type_number_t size);
#endif
