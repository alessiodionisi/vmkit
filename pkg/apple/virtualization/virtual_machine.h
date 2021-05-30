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
void VZVirtualMachine_start(void *ptr, void *queue);
void VZVirtualMachine_pause(void *ptr, void *queue);
void VZVirtualMachine_resume(void *ptr, void *queue);
bool VZVirtualMachine_isSupported();

#endif /* virtual_machine_h */
