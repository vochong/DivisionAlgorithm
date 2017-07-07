# Division Algorithm
Calculates *a/b* to *n* decimal digits of precision given integers *a*, *b* and *n*.
The newest python 3 version is division4.py

There are four main functions in division4.py
* ``` divide(a,b,n) ```
* ``` divide2(a,b,n) ```
* ``` digits(a,b,n1,n2) ```
* ``` digits2(a,b,n1,n2) ```

Next here are some examples of using these functions in division4.py

# Example 1 divide(a,b,n)
In order to use the divide function if *gcd(a,b) = 1* then *gcd(b,10)* must be *1* otherwise an error will occur.  
If *gcd(a,b) > 1* then let *b1 = b/gcd(a,b)* then *gcd(b1,10)* must be *1* otherwise an error message will be output.

The following is an example of a python script that calculates division to *100* decimal places of accuracy:

```python

from division4 import *

a = 1243455434495498469654903493985349490549019934549023567867
b = 54843812848548794568765842358934894313287487348732874378379

k = divide(a,b,100)
print("\n\na =",a,"\nb =",b)
print("\nk =",k)
```

And here is the output:

```python
a = 1243455434495498469654903493985349490549019934549023567867 
b = 54843812848548794568765842358934894313287487348732874378379

k = 0.0226726657012214125309683389217720652317550324675471858158615046233657821289938566786269005119059332
```

In this example the variables *a* and *b* can be integers of arbitrary size however the output *k* will be returned as a string.

# Example 2 divide2(a,b,n)
This is a improved version of the function divide(a,b,n).  Like the function divide, divide2 works in all cases for given integers *a*,*b*, and *n* provided that *gcd(b,10) = 1*. But unlike divide it works in most cases if *gcd(b,10)>1*.  Let *b1 = b/gcd(a,b)*, if *b1* is not divisible by *2<sup>17</sup>* which is *131072* and *b1* is also not divisible by *5<sup>17</sup>* which is *762939453125* then the divide2 will be able to calculate *a/b* to *n* decimal digits of accuracy. 

The following is an example of a script using the function divide2.

```python
from division4 import *

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

# Example 3 digits(a,b,n1,n2)
Similar to the function divide the function digits in order to work correctly requires that if *gcd(a,b) = 1* that *gcd(b,10) = 1*.  If *gcd(a,b)>1* let *b1 = b/gcd(a,b)* it requires that *gcd(b1,10) = 1* otherwise an error message will be output.  The function digits(a,b,n1,n2) calculates the *n1* to *n2* digits in the decimal expansion of the fractional part of *a/b*.  For example the following python script illustrates this:

```python
from division4 import *

a = 632
b = 29

k = divide(a,b,56)
print("\nk =",k)

c = digits(a,b,1,28)
print("\nc =",c)

d = digits(a,b,10,20)
print("\nd =",d)

t = digits(a,b,-9,0)
print("\nt =",t)
```
The following is the output from this script:

```python
k = 21.79310344827586206896551724137931034482758620689655172413

c = 7931034482758620689655172413

d = 27586206896

t = 9655172413
```
*k* is equal to *632/29* to *56* decimal places of accuracy.  The variable *c* is equal to the *1st* to the *28th* digits in the decimal expansion of *632/29* whereas *d* equals the *10th* to the *20th* digits.  The functions digits also allows for the parameters *n1* and *n2* to be *0* or negative integers.  If *0* is used as a value for *n1* or *n2* this denotes the last digit in the decimal expansion of *a/b* before it starts repeating, *-1* denotes the next to last and so on.  Therefore *t* returns the last *10* digits in the decimal expansion of the fractional part of *632/29* before it starts repeating.

The following is another example using large integers that illustrates how in many cases the function digits is much more efficient than other more convential division algorithms:

```python
from division4 import *

a = "142319824384532854895453109433405954879568976589076768767790823490239183527578"
a = a + "95145689454512123432458670489765325768123589231654362785348798142342312568943651"
a = int(a)
b = 2**521-1
n1 = 2**400
n2 = n1+50

print("\na =",a)
print("\nb =",b)
print("\nn1 =",n1)
print("\nn2 =",n2)

k = digits(a,b,n1,n2)
print("\nk =",k)

t = digits(a,b,-19,0)
print("\nt =",t)
```

The output from this example is:

```python
a = 14231982438453285489545310943340595487956897658907676876779082349023918352757895145689454512123432458670489765325768123589231654362785348798142342312568943651

b = 6864797660130609714981900799081393217269435300143305409394463459185543183397656052122559640661454554977296311391480858037121987999716643812574028291115057151

n1 = 2582249878086908589655919172003011874329705792829223512830659356540647622016841194629645353280137831435903171972747493376

n2 = 2582249878086908589655919172003011874329705792829223512830659356540647622016841194629645353280137831435903171972747493426

k = 530885742165702296952553483592649507952301281362994

t = 40916999596875838501
```

First notice that in this example *a* was first defined to be a string by adding two strings together and then converting that string to an integer.  This wasn't necessary, but was done in this example so that it would be easier to read and so hopefully there won't be any problems if someone copies and pastes this example.  The variable *a* is just some random *158* digit number and *b* is a *521* bit or *157* digit prime number.  This program outputs *k* which is the *n1* to *n2* digits in the decimal expansion of *a/b* and *t* which is the last *20* digits in the decimal expansion of *a/b* before it starts repeating.  Using more convential floating point division algorithms calculating either of these two values would be impossible.  For example to calculate the value of *k* using other more commonly used floating point division algorithms, one would have to calculate *a/b* to *n2* digits of accuracy.  Given how large *n2* is in this example, this would be impossible.

# Example 4 digits2(a,b,n1,n2)
This is an improved version of the function digits(a,b,n1,n2) and unlike the other three functions it works for all possible integer values of *a* and *b*.  The function digits2 returns the 2-tuple *(c , s)* where *c* is integer that indicates how many digits in the decimal expansion of *a/b* have to be skipped until it starts repeating and *s* is a string that is equal to the *n1* to *n2* digits on the repeating part of the decimal expansion of *a/b*.

The following is a example of a script using the function digits2(a,b,n1,n2):

```python

from division4 import *

a = 29
b = 124

print("\na =",a)
print("b =",b)

d = divide2(a,b,32)
print("\nd =",d)

k = digits2(a,b,1,10)
print("\nk =",k)

t = digits2(a,b,-19,0)
print("\nt =",t)
```

The following is the output from this script:

```python

a = 29
b = 124

d = 0.23387096774193548387096774193548

k = (2, '3870967741')

t = (2, '93548387096774193548')
```

The variable *d* is equal to *a/b* to 32 decimal digits of accuracy, k is a 2-tuple of an integer 2 and a string.  The value 2 indicates that 2 places have to be skipped until the repeating part of the decimal expansion of *a/b* begins and the string represents the first 10 digits in that decimal expansion.  Similarly for the variable *t* the string represent the last 20 digits in that decimal expansion.


