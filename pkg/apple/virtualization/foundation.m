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

#import "foundation.h"

const char *NSError_localizedDescription(void *ptr) {
  return [[(NSError *)ptr localizedDescription] UTF8String];
}

void *NSMutableArray_arrayWithCapacity(unsigned long capacity) {
  return [NSMutableArray arrayWithCapacity:(NSUInteger)capacity];
}

void NSMutableArray_addObject(void *ptr, void *object) {
  [(NSMutableArray *)ptr addObject:object];
}

void *newDispatchQueue(const char *label) {
  return dispatch_queue_create(label, DISPATCH_QUEUE_SERIAL);
}
