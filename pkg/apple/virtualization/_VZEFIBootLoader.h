#import <Virtualization/VZDefines.h>

NS_ASSUME_NONNULL_BEGIN

VZ_EXPORT
@interface _VZEFIBootLoader : VZBootLoader

- (instancetype)init NS_DESIGNATED_INITIALIZER;

@property(nullable, copy) NSURL *efiURL;
@property(nullable, copy) _VZEFIVariableStore *variableStore;

@end

NS_ASSUME_NONNULL_END
