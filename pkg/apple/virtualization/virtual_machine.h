// Spin up Linux VMs with QEMU and Apple virtualization framework
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

#ifndef virtual_machine_h
#define virtual_machine_h

#import <Foundation/Foundation.h>
#import <Virtualization/Virtualization.h>

void *VZVirtualMachine_init(void *configuration, void *queue);
int VZVirtualMachine_state(void *ptr, void *queue);
bool VZVirtualMachine_canStart(void *ptr, void *queue);
bool VZVirtualMachine_canPause(void *ptr, void *queue);
bool VZVirtualMachine_canResume(void *ptr, void *queue);
bool VZVirtualMachine_canRequestStop(void *ptr, void *queue);
void VZVirtualMachine_start(void *ptr, void *queue, const char *handlerID);
void VZVirtualMachine_pause(void *ptr, void *queue);
void VZVirtualMachine_resume(void *ptr, void *queue);
bool VZVirtualMachine_isSupported();

#endif /* virtual_machine_h */
