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
