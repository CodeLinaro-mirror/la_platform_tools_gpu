LOCAL_PATH := $(call my-dir)

include $(CLEAR_VARS)

# Use the pre-built protobuf library.
SRC_FILE_LIST          := $(LOCAL_PATH)/../../../../../external/protobuf/prebuilts/$(TARGET_ARCH)/libprotobuf-cpp-2.3.0-lite.a
LOCAL_MODULE           := libprotobuf
LOCAL_SRC_FILES        := $(SRC_FILE_LIST:$(LOCAL_PATH)/%=%)
include $(PREBUILT_STATIC_LIBRARY)

include $(CLEAR_VARS)

SRC_FILE_LIST          := \
    $(wildcard $(LOCAL_PATH)/../../src/gfxspy/EGL/*.cpp) \
    $(wildcard $(LOCAL_PATH)/../../src/gfxspy/*.cpp)
LOCAL_MODULE           := gfxspy
LOCAL_SRC_FILES        := $(SRC_FILE_LIST:$(LOCAL_PATH)/%=%)
LOCAL_CXXFLAGS         := \
    $(libc_common_cflags) \
    -std=c++11 \
    -DEGL_EGLEXT_PROTOTYPES \
    -DGL_GLEXT_PROTOTYPES \
    -DTARGET_OS_ANDROID \
    -DANDROID_NDK \
    -DGOOGLE_PROTOBUF_NO_RTTI \
    -DHAVE_SYS_UIO_H \
    -DHAVE_PTHREADS \
    -DGENERATE_API_WRAPPERS \
    -DLOAD_HOOKS_AT_STARTUP \
    -D__NDK_FPABI_MATH__= \
    -DOS_PATH_SEPARATOR=\'/\' \
    -Ofast \
    -marm \
    -mfloat-abi=softfp \
    -fno-strict-aliasing \
    -funswitch-loops \
    -finline-limit=100 \
    -fno-omit-frame-pointer \
    -ffunction-sections \
    -fpic \
    -fstack-protector \
    -funwind-tables \
    -fno-short-enums \
    -fno-exceptions
LOCAL_CFLAGS         := \
    $(libc_common_cflags)
LOCAL_C_INCLUDES       := \
    $(LOCAL_PATH)/../../src/gfxspy/ \
    $(LOCAL_PATH)/../../../../../frameworks/native/opengl/include \
    $(LOCAL_PATH)/../../../../../frameworks/native/opengl/libs \
    $(LOCAL_PATH)/../../../../../external/protobuf/src \
    $(LOCAL_PATH)/../../../../../system/core/include \
    $(LOCAL_PATH)/../../../../../bionic/libc/include \
    $(LOCAL_PATH)/../../../../../bionic/libc/private \
    /usr/local/include
LOCAL_LDLIBS           := -lc -lm -llog
LOCAL_STATIC_LIBRARIES := libprotobuf

include $(BUILD_SHARED_LIBRARY)

