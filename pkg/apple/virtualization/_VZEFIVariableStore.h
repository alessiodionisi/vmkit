#import <Virtualization/VZDefines.h>

NS_ASSUME_NONNULL_BEGIN

VZ_EXPORT
@interface _VZEFIVariableStore : NSObject <NSCopying>

- (instancetype)initWithURL:(NSURL *)url
                      error:(NSError *_Nullable *)error
    NS_DESIGNATED_INITIALIZER;

@end

NS_ASSUME_NONNULL_END
