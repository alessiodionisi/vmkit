#import "virtual_machine.h"

void *VZVirtualMachine_init(void *configuration, void *queue) {
  return [[VZVirtualMachine alloc]
      initWithConfiguration:(VZVirtualMachineConfiguration *)configuration
                      queue:(dispatch_queue_t)queue];
}

bool VZVirtualMachine_canStart(void *ptr, void *queue) {
  __block bool canStart;

  dispatch_sync((dispatch_queue_t)queue, ^{
    canStart = [(VZVirtualMachine *)ptr canStart];
  });

  return canStart;
}

void VZVirtualMachine_startWithCompletionHandler(void *ptr, void *queue) {
  dispatch_sync((dispatch_queue_t)queue, ^{
    [(VZVirtualMachine *)ptr startWithCompletionHandler:^(NSError *error){

    }];
  });
}
