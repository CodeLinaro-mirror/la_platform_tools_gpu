APP_ABI := armeabi-v7a arm64-v8a
APP_PLATFORM := android-21
APP_STL := stlport_static
NDK_TOOLCHAIN_VERSION := clang
# libprotobuf was compiled with -fno-rtti, so we must do the same for gfxspy.
# This can not go in Android.mk because ndk-build inserts -frtti after
# $LOCAL_CPPFLAGS.  ndk-build does not insert any options after $APP_CPPFLAGS.
APP_CPPFLAGS := -fno-rtti
