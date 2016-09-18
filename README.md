# DivisionAlgorithm
Python 3 and Go 1.7 implementations to calculate a/b given two integers a and b to n decimal digits of precision

The main two functions are divide(a,b,n) which calculates a/b to n decimal digits of precision and digits(a,b,n1,n2) which 
calculates the n1 to n2 digits in the decimal expansion of the fractional part of a/b.  In the Python implementation the integers
a and b can be of arbitrary size.  In the Go version the type int is used which is an 8 byte size integer meaning that a and b 
should be less than 2^31.5 (since one bit is used for the sign) to work properly.
