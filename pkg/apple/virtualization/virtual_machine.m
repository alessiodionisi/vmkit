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
