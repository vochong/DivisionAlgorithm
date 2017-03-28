# Division Algorithm
Calculates a/b to n decimal digits of precision given integers a, b and n.
There are golang versions and python 3 versions. 
The newest python version is division3.py

There are four main functions in division3.py
* ``` divide(a,b,n) ```
* ``` divide2(a,b,n) ```
* ``` digits(a,b,n1,n2) ```
* ``` digits2(a,b,n1,n2) ```

Next here are some examples of using these functions in division3.py

# Example 1 divide(a,b,n)
In order to use the divide function if gcd(a,b) = 1 then gcd(b,10) must be 1 otherwise an error will occur.  
If gcd(a,b) > 1 then let b1 = b/gcd(a,b) then gcd(b1,10) must be 1 otherwise an error message will be output.

The following is an example of a python script that calculates division to 100 decimal place of accuracy:

```python

from division3 import *

a = 1243455434495498469654903493985349490549019934549023567867
b = 54843812848548794568765842358934894313287487348732874378379

y = divide(a,b,100)
print("\n\na =",a,"\nb =",b)
print("\na/b =",y)
```

And here is the output:

```python
a = 1243455434495498469654903493985349490549019934549023567867 
b = 54843812848548794568765842358934894313287487348732874378379

a/b = 0.0226726657012214125309683389217720652317550324675471858158615046233657821289938566786269005119059332
```

In this example the variables a and b can be integers of arbitrary size however the output y will be returned as a string.

# Example 2 divide2(a,b,n)
This is a improved version of the function divide(a,b,n).  Like the function divide, divide2 works in all cases for given integers a,b, and n provided that gcd(b,10) = 1. But unlike divide it works in most cases if gcd(b,10)>1.  Let b1 = b/gcd(a,b),
if b1 is not divisible by 2^17 which is 131072 and b1 is also not divisible by 5^17 which is 762939453125 then the divide2 will be able to calculate a/b to n decimal digits of accuracy. 

The following is an example of a script using the function divide2.

```python
from division3 import *

a = 34329804329054324324349865294321238390324609435465789876543958423904539076023134295647
t1 = 2970524338954839323842046534309342412487956546707676875645324365749
t2 = 2**11
t3 = 5**16
b = t1*t2*t3
y = divide2(a,b,180)

print("\n\na =",a,"\nb =",b)
print("\na/b =",y)
```

And this is the output for example 2.

```python
a = 34329804329054324324349865294321238390324609435465789876543958423904539076023134295647 
b = 928288855923387288700639541971669503902486420846149023639163864296562500000000000

a/b = 36981.81240677050769387854186412487202403005977623028962464219192411000956034137055640905286098818
59515314907258766772144351232503401122044415586143731297419997286703268341777442322425678656601652434103
```




