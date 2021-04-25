#ifndef serial_port_h
#define serial_port_h

#import <Foundation/Foundation.h>
#import <Virtualization/Virtualization.h>

void *VZVirtioConsoleDeviceSerialPortConfiguration_init(void *attachment);
void *VZFileHandleSerialPortAttachment_init(int read, int write);

#endif /* serial_port_h */
