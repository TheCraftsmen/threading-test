from ctypes import cdll
lib = cdll.LoadLibrary('./shfp.so')
lib.add()