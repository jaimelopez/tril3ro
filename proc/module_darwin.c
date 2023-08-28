#include "module_darwin.h"
#include "../mem/mem_darwin.c"

struct Module *get_modules(pid_t pid, uint32_t *count)
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
        uintptr_t data = read_process_memory(task, dyld_info.all_image_info_addr, &size_infos);
        struct dyld_all_image_infos *infos = (struct dyld_all_image_infos *)data;

        mach_msg_type_number_t size_count = sizeof(struct dyld_image_info) * infos->infoArrayCount;
        uintptr_t info_addr = read_process_memory(task, (mach_vm_address_t)infos->infoArray, &size_count);
        struct dyld_image_info *info = (struct dyld_image_info *)info_addr;

        *count = infos->infoArrayCount;

        struct Module *items = malloc(infos->infoArrayCount * sizeof(struct Module));

        for (int i = 0; i < infos->infoArrayCount; i++)
        {
                mach_msg_type_number_t size_path = PATH_MAX;
                uintptr_t fpath_addr = read_process_memory(task, (mach_vm_address_t)info[i].imageFilePath, &size_path);
                const char *fpath = (const char *)fpath_addr;

                if (!fpath)
                {
                        continue;
                }

                struct stat st;
                stat(fpath, &st);

                items[i].addr = (mach_vm_address_t)info[i].imageLoadAddress;
                items[i].module = fpath;
                items[i].size = st.st_size;
        }

        return items;
}
