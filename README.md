# Division Algorithm
Calculates a/b to n decimal digits of precision given integers a, b and n.

There are golang versions and python 3 versions. The newest python version is division3.py

There are four main functions in division3.py
* ``` divide(a,b,n) ```
* ``` divide2(a,b,n) ```
* ``` digits(a,b,n1,n2) ```
* ``` digits2(a,b,n1,n2) ```

Next here are some examples of using these functions in division3.py

# Example 1 divide(a,b,n)
In order to use the divide function if gcd(a,b) = 1 then gcd(b,10) must be 1 otherwise an error will occur.  
If gcd(a,b) > 1 then let b1 = b/gcd(a,b) then gcd(b1,10) must be 1 otherwise an error message will be output.

The following is an example of a python script that calculate division to 100 decimal place of accuracy:

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

