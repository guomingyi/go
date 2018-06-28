#!/bin/bash

path=$1
ACTION=$2

BL=bootloader-v1.0.bin
FW=firmware-v1.0.bin
MT=metadata-v1.0.bin
FTM=ftm.bin
FTM_M=metadata-ftm.bin

BL_ADDR=0x08000000
MT_ADDR=0x08008000
FW_ADDR=0x08010000

function device_reset()
{
  echo -e "\033[32m\n st-flash reset. \033[0m"
  st-flash reset
}

if [ "y$ACTION" = "yb" ] ;then
  echo -e "\033[32m\n st-flash write $path/$BL $BL_ADDR. \033[0m"
  st-flash write $path/$BL $BL_ADDR
  device_reset
  exit 0
fi

if [ "y$ACTION" = "yf" ]; then
  echo -e "\033[32m\n st-flash write  $path/$FW $FW_ADDR. \033[0m"
  st-flash write  $path/$FW $FW_ADDR
  device_reset
  exit 0
fi

if [ "y$ACTION" = "ym" ]; then
  echo -e "\033[32m\n st-flash write  $path/$MT $MT_ADDR. \033[0m"
  st-flash write  $path/$MT $MT_ADDR
  device_reset
  exit 0
fi

if [ "y$ACTION" = "yftm-m" ]; then
  echo -e "\033[32m\n 3 ==> st-flash write  $path/$FTM_M $MT_ADDR. \033[0m"
  st-flash write  $path/$FTM_M $MT_ADDR
  device_reset
  exit 0
fi

if [ "y$ACTION" = "yftm-f" ]; then
  echo -e "\033[32m\n 1 ==> st-flash write  $path/$FTM $FW_ADDR. \033[0m"
  st-flash write  $path/$FTM $FW_ADDR
  device_reset
  exit 0
fi



