#ifndef virtual_machine_h
#define virtual_machine_h

#import <Foundation/Foundation.h>
#import <Virtualization/Virtualization.h>

void *VZVirtualMachine_init(void *configuration, void *queue);
bool VZVirtualMachine_canStart(void *ptr, void *queue);
void VZVirtualMachine_startWithCompletionHandler(void *ptr, void *queue);

#endif /* virtual_machine_h */
