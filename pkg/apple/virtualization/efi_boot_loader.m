#import "efi_boot_loader.h"

void *VZEFIVariableStore_init(const char *url, void **error) {
  NSString *string = [NSString stringWithUTF8String:url];
  NSURL *parsedURL = [NSURL fileURLWithPath:string];

  return [[_VZEFIVariableStore alloc] initWithURL:parsedURL
                                            error:(NSError *_Nullable *)error];
}

void *VZEFIBootLoader_init() { return [[_VZEFIBootLoader alloc] init]; }

void VZEFIBootLoader_setEfiURL(void *ptr, const char *efiURL) {
  NSString *string = [NSString stringWithUTF8String:efiURL];
  NSURL *url = [NSURL fileURLWithPath:string];

  [(_VZEFIBootLoader *)ptr setEfiURL:url];
}

void VZEFIBootLoader_setVariableStore(void *ptr, void *variableStore) {
  [(_VZEFIBootLoader *)ptr
      setVariableStore:(_VZEFIVariableStore *)variableStore];
}
