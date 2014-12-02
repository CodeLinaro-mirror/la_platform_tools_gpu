APP_ABI := armeabi-v7a
APP_CPPFLAGS += -std=c++11
APP_PLATFORM := android-21
APP_STL := stlport_static
NDK_TOOLCHAIN_VERSION := clang
APP_OPTIM := debug
# libprotobuf was compiled with -fno-rtti, so we must do the same for gfxspy.
# This can not go in Android.mk because ndk-build inserts -frtti after
# $LOCAL_CPPFLAGS.  ndk-build does not insert any options after $APP_CPPFLAGS.
APP_CPPFLAGS := -fno-rtti
