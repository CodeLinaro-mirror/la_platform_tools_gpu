LOCAL_PATH := $(call my-dir)

include $(CLEAR_VARS)

# Use the pre-built protobuf library.
SRC_FILE_LIST          := $(LOCAL_PATH)/../../../../../../external/protobuf/prebuilts/$(TARGET_ARCH)/libprotobuf-cpp-2.3.0-lite.a
LOCAL_MODULE           := libprotobuf
LOCAL_SRC_FILES        := $(SRC_FILE_LIST:$(LOCAL_PATH)/%=%)
include $(PREBUILT_STATIC_LIBRARY)

include $(CLEAR_VARS)

SRC_FILE_LIST          := $(wildcard $(LOCAL_PATH)/../../src/*.cpp)
LOCAL_MODULE           := gfxspy
LOCAL_SRC_FILES        := $(SRC_FILE_LIST:$(LOCAL_PATH)/%=%)
LOCAL_CXXFLAGS         := \
    -std=c++11 \
    -DEGL_EGLEXT_PROTOTYPES \
    -DGL_GLEXT_PROTOTYPES \
    -DTARGET_OS_ANDROID \
    -DANDROID_NDK \
    -DGLTRACE_GENERATE_API_WRAPPERS \
    -DGLTRACE_DLOPEN_INTERCEPTION \
    -DGOOGLE_PROTOBUF_NO_RTTI \
    -DHAVE_SYS_UIO_H \
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
    -fno-exceptions \
    -Wno-deprecated-register

#############
# Optional flags for debugging:
#    -DGLTRACE_DISABLE_SERVER \
#    -DGLTRACE_PRINT_EGL_CALLS \
#    -DGLTRACE_SHOULDNT_LOAD_HOOKS_AT_STARTUP \

LOCAL_C_INCLUDES       := \
    $(LOCAL_PATH)/../../src/ \
    $(LOCAL_PATH)/../../../../../../frameworks/native/opengl/include \
    $(LOCAL_PATH)/../../../../../../external/protobuf/src \
    /usr/local/include
LOCAL_LDLIBS           := -lc -lm -llog
LOCAL_STATIC_LIBRARIES := libprotobuf

include $(BUILD_SHARED_LIBRARY)

