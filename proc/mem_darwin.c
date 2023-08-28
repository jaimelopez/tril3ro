#include "mem_darwin.h"

uintptr_t read_process_memory(task_t task, mach_vm_address_t addr, mach_msg_type_number_t *size)
{
    mach_msg_type_number_t count = (mach_msg_type_number_t)*size;
    vm_offset_t data;

    if (vm_read(task, addr, *size, &data, &count) != KERN_SUCCESS)
    {
        return 0;
    }

    return (uintptr_t)data;
}

void read_process_memory_bytes(task_t task, mach_vm_address_t addr, void *buffer, mach_msg_type_number_t *size)
{
    vm_size_t count = (vm_size_t)size;
    vm_read_overwrite(task, addr, *size, (vm_address_t)buffer, &count);
}

bool write_process_memory(task_t task, mach_vm_address_t addr, vm_offset_t data, mach_msg_type_number_t size)
{
    if (addr == 0)
    {
        return false;
    }

    return (vm_write(task, addr, data, size) == KERN_SUCCESS);
}
