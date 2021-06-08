#import "virtual_machine.h"
#import "_cgo_export.h"

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

bool VZVirtualMachine_canPause(void *ptr, void *queue) {
  __block bool canPause;

  dispatch_sync((dispatch_queue_t)queue, ^{
    canPause = [(VZVirtualMachine *)ptr canPause];
  });

  return canPause;
}

bool VZVirtualMachine_canResume(void *ptr, void *queue) {
  __block bool canResume;

  dispatch_sync((dispatch_queue_t)queue, ^{
    canResume = [(VZVirtualMachine *)ptr canResume];
  });

  return canResume;
}

bool VZVirtualMachine_canRequestStop(void *ptr, void *queue) {
  __block bool canRequestStop;

  dispatch_sync((dispatch_queue_t)queue, ^{
    canRequestStop = [(VZVirtualMachine *)ptr canRequestStop];
  });

  return canRequestStop;
}

int VZVirtualMachine_state(void *ptr, void *queue) {
  __block int state;

  dispatch_sync((dispatch_queue_t)queue, ^{
    state = [(VZVirtualMachine *)ptr state];
  });

  return state;
}

void VZVirtualMachine_start(void *ptr, void *queue, const char *handlerID) {
  dispatch_sync((dispatch_queue_t)queue, ^{
    [(VZVirtualMachine *)ptr
        startWithCompletionHandler:Block_copy(^(NSError *error) {
          startErrorHandler(error, (char *)handlerID);
        })];
  });
}

void VZVirtualMachine_pause(void *ptr, void *queue) {
  dispatch_sync((dispatch_queue_t)queue, ^{
    [(VZVirtualMachine *)ptr pauseWithCompletionHandler:^(NSError *error) {
      NSLog(@"%@", error);
    }];
  });
}

void VZVirtualMachine_resume(void *ptr, void *queue) {
  dispatch_sync((dispatch_queue_t)queue, ^{
    [(VZVirtualMachine *)ptr resumeWithCompletionHandler:^(NSError *error) {
      NSLog(@"%@", error);
    }];
  });
}

bool VZVirtualMachine_isSupported() { return [VZVirtualMachine isSupported]; }
