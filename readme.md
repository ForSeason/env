# How to use

Set `.env` file in the root of your workspace. This file will be loaded while this package is imported.

If initialization fails, the package will print error in log and **clear all data that has been loaded in memory** .

Initializer use `=` as seprator between key and value. That is to say, part before first `=` is key, and the rest part is value. It is ok to use empty value.

Initializer will ignore empty line.
