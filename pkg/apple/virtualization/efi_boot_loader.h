#ifndef efi_boot_loader_h
#define efi_boot_loader_h

#import <Virtualization/Virtualization.h>

#import "_VZEFIVariableStore.h"

#import "_VZEFIBootLoader.h"

void *VZEFIVariableStore_init(const char *url, void **error);

void *VZEFIBootLoader_init();
void VZEFIBootLoader_setEfiURL(void *ptr, const char *efiURL);
void VZEFIBootLoader_setVariableStore(void *ptr, void *variableStore);

#endif /* efi_boot_loader_h */
