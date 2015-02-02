# Copyright 2014, The Android Open Source Project
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

LOCAL_PATH := $(call my-dir)
GTEST_PATH := $(LOCAL_PATH)/../../../../../../external/gtest
GMOCK_PATH := $(LOCAL_PATH)/../../../../../../external/gmock

# Compile the replay daemon into a static library
include $(CLEAR_VARS)
# Src files used by the replay daemon expect Main.cpp to avoid conflicts with multiply main functions
SRC_FILE_LIST          := $(filter-out $(LOCAL_PATH)/../../src/Main.cpp, $(wildcard $(LOCAL_PATH)/../../src/*.cpp))

LOCAL_MODULE           := caze_static
LOCAL_SRC_FILES        := $(SRC_FILE_LIST:$(LOCAL_PATH)/%=%)

LOCAL_CFLAGS           := -std=c++11 -DEGL_EGLEXT_PROTOTYPES -DGL_GLEXT_PROTOTYPES -DTARGET_OS_ANDROID -Ofast
LOCAL_C_INCLUDES       := /usr/local/include
include $(BUILD_STATIC_LIBRARY)

# Compile the replay daemon into a shared library for the apk creation
include $(CLEAR_VARS)
# Only Main.cpp have to be compiled because everything else came from the caze static library
SRC_FILE_LIST          := $(LOCAL_PATH)/../../src/Main.cpp

LOCAL_MODULE           := caze
LOCAL_SRC_FILES        := $(SRC_FILE_LIST:$(LOCAL_PATH)/%=%)

LOCAL_CFLAGS           := -std=c++11 -DEGL_EGLEXT_PROTOTYPES -DGL_GLEXT_PROTOTYPES -DTARGET_OS_ANDROID -Ofast
LOCAL_C_INCLUDES       := /usr/local/include
LOCAL_LDLIBS           := -lEGL -lGLESv1_CM -lGLESv3 -llog -landroid -lz
LOCAL_STATIC_LIBRARIES := android_native_app_glue caze_static
include $(BUILD_SHARED_LIBRARY)

# Compile gtest into a static library
include $(CLEAR_VARS)
SRC_FILE_LIST          := $(GTEST_PATH)/src/gtest-all.cc $(GTEST_PATH)/src/gtest_main.cc

LOCAL_MODULE           := gtest
LOCAL_SRC_FILES        := $(SRC_FILE_LIST:$(LOCAL_PATH)/%=%)

LOCAL_CFLAGS           := -std=c++11
LOCAL_C_INCLUDES       := $(GTEST_PATH) $(GTEST_PATH)/include
include $(BUILD_STATIC_LIBRARY)

# Compile gmock into a static library
include $(CLEAR_VARS)
SRC_FILE_LIST          := $(GMOCK_PATH)/src/gmock-all.cc

LOCAL_MODULE           := gmock
LOCAL_SRC_FILES        := $(SRC_FILE_LIST:$(LOCAL_PATH)/%=%)

LOCAL_CFLAGS           := -std=c++11
LOCAL_C_INCLUDES       := $(GMOCK_PATH) $(GMOCK_PATH)/include $(GTEST_PATH)/include
include $(BUILD_STATIC_LIBRARY)

# Compile the test files into an Android executable
include $(CLEAR_VARS)
# Only the test files have to be compiled here because everything else came from static libraries
SRC_FILE_LIST          := $(wildcard $(LOCAL_PATH)/../../test/*.cpp)

LOCAL_MODULE           := caze_test
LOCAL_SRC_FILES        := $(SRC_FILE_LIST:$(LOCAL_PATH)/%=%)

LOCAL_CFLAGS           := -std=c++11 -DEGL_EGLEXT_PROTOTYPES -DGL_GLEXT_PROTOTYPES -DTARGET_OS_ANDROID
LOCAL_C_INCLUDES       := /usr/local/include $(LOCAL_PATH)/../../src $(GTEST_PATH)/include $(GMOCK_PATH)/include
LOCAL_LDLIBS           := -lEGL -lGLESv1_CM -lGLESv3 -llog -landroid -lz
LOCAL_STATIC_LIBRARIES := gtest gmock caze_static
include $(BUILD_EXECUTABLE)

$(call import-module,android/native_app_glue)
