LOCAL_PATH := $(call my-dir)

include $(CLEAR_VARS)

# Use the pre-built protobuf library.
SRC_FILE_LIST          := $(LOCAL_PATH)/../../../../../external/protobuf/prebuilts/arm/libprotobuf-cpp-2.3.0-lite.a
LOCAL_MODULE           := libprotobuf
LOCAL_SRC_FILES        := $(SRC_FILE_LIST:$(LOCAL_PATH)/%=%)
include $(PREBUILT_STATIC_LIBRARY)

include $(CLEAR_VARS)

# TODO: Change this to *.cpp once all the files are ready to compile.
SRC_FILE_LIST          := $(wildcard $(LOCAL_PATH)/../../src/gfxspy/gltrace.pb.cpp)

LOCAL_MODULE           := gfxspy
LOCAL_SRC_FILES        := $(SRC_FILE_LIST:$(LOCAL_PATH)/%=%)

LOCAL_CPPFLAGS         := -DEGL_EGLEXT_PROTOTYPES -DGL_GLEXT_PROTOTYPES -DTARGET_OS_ANDROID -DANDROID_NDK -Ofast -DGOOGLE_PROTOBUF_NO_RTTI -DHAVE_SYS_UIO_H -DHAVE_PTHREADS -DGENERATE_API_WRAPPERS -DLOAD_HOOKS_AT_STARTUP -O0 -marm -fno-strict-aliasing -funswitch-loops -finline-limit=100 -fno-omit-frame-pointer -ffunction-sections -mfloat-abi=softfp -fpic -fstack-protector -funwind-tables -fno-short-enums -fno-exceptions

LOCAL_C_INCLUDES       := /usr/local/include $(LOCAL_PATH)/../../../../../external/protobuf/src
LOCAL_LDLIBS           := -lc -lm -llog
LOCAL_STATIC_LIBRARIES := libprotobuf libstlport_static
LOCAL_SHARED_LIBRARIES := -lgcc

include $(BUILD_SHARED_LIBRARY)
