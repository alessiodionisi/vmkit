// Virtual Machine manager that supports QEMU and Apple virtualization framework on macOS
// Copyright (C) 2021 VMKit Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
