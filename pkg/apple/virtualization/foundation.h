#ifndef foundation_h
#define foundation_h

#import <Foundation/Foundation.h>

const char *NSError_localizedDescription(void *ptr);
void *NSMutableArray_arrayWithCapacity(unsigned long capacity);
void NSMutableArray_addObject(void *ptr, void *object);
void *newDispatchQueue(const char *label);

#endif /* foundation_h */
