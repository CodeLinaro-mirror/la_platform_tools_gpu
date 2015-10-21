#!/bin/bash

SDK_PATH=gapid
VERSION=${SDK_PATH}_r01
INPUT=bin
OUTPUT=temp-sdk
LINUX=${VERSION}_linux

# First we build the basic set of files used in all sdk packages
BASE_SDK=${OUTPUT}/base/${SDK_PATH}
BASE_OUT=${OUTPUT}/base/${SDK_PATH}
BASE_ANDROID=${BASE_OUT}/android
mkdir -p ${BASE_ANDROID}
cp  ${INPUT}/pkginfo.apk ${BASE_ANDROID}/
mkdir -p ${BASE_ANDROID}/armeabi-v7a
cp  ${INPUT}/android-arm/release/libgapii.so ${BASE_ANDROID}/armeabi-v7a/
cp  ${INPUT}/android-arm/release/gapir.apk ${BASE_ANDROID}/armeabi-v7a/
mkdir -p ${BASE_ANDROID}/arm64-v8a
cp  ${INPUT}/android-arm64/release/libgapii.so ${BASE_ANDROID}/arm64-v8a/
cp  ${INPUT}/android-arm64/release/gapir.apk ${BASE_ANDROID}/arm64-v8a/

# Now we build the os specific versions
LINUX_SDK=${OUTPUT}/${LINUX}
LINUX_OUT=${OUTPUT}/${LINUX}/${SDK_PATH}
mkdir -p ${LINUX_OUT}
cp -r ${BASE_SDK} ${LINUX_SDK}
mkdir -p ${LINUX_OUT}/linux/x86_64
cp  ${INPUT}/gapis ${LINUX_OUT}/linux/x86_64/
cp  ${INPUT}/gapit ${LINUX_OUT}/linux/x86_64/
cp  ${INPUT}/linux-X86_64/release/gapir ${LINUX_OUT}/linux/x86_64/
cd ${OUTPUT}/${LINUX} && zip -9 -r ${LINUX}.zip ${SDK_PATH}
