#include "module_darwin.h"

struct Module* getModules(pid_t pid, uint32_t* count)
{
        task_t task;
        task_for_pid(mach_task_self(), pid, &task);

        struct task_dyld_info dyld_info;
        mach_msg_type_number_t info_count = TASK_DYLD_INFO_COUNT;

        if (task_info(task, TASK_DYLD_INFO, (task_info_t)&dyld_info, &info_count) != KERN_SUCCESS)
        {
                return (struct Module *)(NULL);
        }

        mach_msg_type_number_t size_infos = sizeof(struct dyld_all_image_infos);
        uint8_t *data = readProcessMemory(pid, dyld_info.all_image_info_addr, &size_infos);
        struct dyld_all_image_infos *infos = (struct dyld_all_image_infos *)data;

        mach_msg_type_number_t size_count = sizeof(struct dyld_image_info) * infos->infoArrayCount;
        uint8_t *info_addr = readProcessMemory(pid, (mach_vm_address_t)infos->infoArray, &size_count);
        struct dyld_image_info *info = (struct dyld_image_info *)info_addr;

        *count = infos->infoArrayCount;

        struct Module * items = malloc(infos->infoArrayCount * sizeof(struct Module));

        for (int i = 0; i < infos->infoArrayCount; i++)
        {
                mach_msg_type_number_t size_header = sizeof(struct mach_header);
                uint8_t *header_addr = readProcessMemory(pid, (mach_vm_address_t)info[i].imageLoadAddress, &size_header);
                struct mach_header *header = (struct mach_header *)header_addr;

                mach_msg_type_number_t size_path = PATH_MAX;
                uint8_t *fpath_addr = readProcessMemory(pid, (mach_vm_address_t)info[i].imageFilePath, &size_path);
                const char* fpath = (const char*)fpath_addr;

                if (!fpath)
                {
                        continue;
                }

                items[i].addr = (mach_vm_address_t)info[i].imageLoadAddress;
                items[i].module = fpath;
                items[i].size = header->sizeofcmds;
        }

        return items;
}

unsigned char* readProcessMemory(int pid, mach_vm_address_t addr, mach_msg_type_number_t *size)
{
        task_t task;
        task_for_pid(mach_task_self(), pid, &task);

        mach_msg_type_number_t count = (mach_msg_type_number_t)*size;
        vm_offset_t data;

        if (vm_read(task, addr, *size, &data, &count) != KERN_SUCCESS)
        {
                return NULL;
        }

        return (unsigned char *)data;
}