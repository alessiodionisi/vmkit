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

#import "storage_device.h"

void *VZDiskImageStorageDeviceAttachment_init(const char *diskImageURL,
                                              bool readOnly, void **error) {
  NSString *string = [NSString stringWithUTF8String:diskImageURL];
  NSURL *url = [NSURL fileURLWithPath:string];

  return [[VZDiskImageStorageDeviceAttachment alloc]
      initWithURL:url
         readOnly:readOnly
            error:(NSError *_Nullable *)error];
}

void *VZVirtioBlockDeviceConfiguration_init(void *attachment) {
  return [[VZVirtioBlockDeviceConfiguration alloc]
      initWithAttachment:(VZStorageDeviceAttachment *)attachment];
}
