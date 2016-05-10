# Strace Analyze

## About

Strace Analyze takes in an absolute filepath of an strace output file as a command line argument and logs the frequency that specfic system calls appear. This is useful in figuring out whether executables are partaking in malicious actions because it aids in showing whether unnecessary system calls are being made.

## Usage

./strace-analyze PATH/TO/STRACEFILE


