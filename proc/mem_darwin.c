#include "mem_darwin.h"

uintptr_t readProcessMemory(int pid, mach_vm_address_t addr, mach_msg_type_number_t *size)
{
    task_t task;
    task_for_pid(mach_task_self(), pid, &task);

    mach_msg_type_number_t count = (mach_msg_type_number_t)*size;
    vm_offset_t data;

    if (vm_read(task, addr, *size, &data, &count) != KERN_SUCCESS)
    {
        return 0;
    }

    return (uintptr_t)data;
}

bool writeProcessMemory(int pid, mach_vm_address_t addr, vm_offset_t data, mach_msg_type_number_t size)
{
    if (addr == 0)
    {
        return false;
    }

    task_t task;
    task_for_pid(mach_task_self(), pid, &task);

    return (vm_write(task, addr, data, size) == KERN_SUCCESS);
}