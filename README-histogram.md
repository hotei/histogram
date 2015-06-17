<center>
# histogram
</center>

Comments can be sent to <hotei1352@gmail.com>

## OVERVIEW

Simple histogram program

Output looks like this:

```
Histogram: Test
Bin[   0 -   50] ******************************************************************************************           485
Bin[  50 -  100] ******************************************************************************************           488
Bin[ 100 -  150] ******************************************************************************************           487
Bin[ 150 -  200] **************************************************************************************               466
Bin[ 200 -  250] ***************************************************************************************************  535
Bin[ 250 -  300] *******************************************************************************************          491
Bin[ 300 -  350] ***********************************************************************************************      511
Bin[ 350 -  400] *******************************************************************************************          494
Bin[ 400 -  450] *************************************************************************************************    521
Bin[ 450 -  500] *************************************************************************************************    524
Bin[ 500 -  550] ******************************************************************************************           484
Bin[ 550 -  600] *********************************************************************************************        504
Bin[ 600 -  650] **************************************************************************************************** 537
Bin[ 650 -  700] *****************************************************************************************            478
Bin[ 700 -  750] ***************************************************************************************              470
Bin[ 750 -  800] ************************************************************************************************     516
Bin[ 800 -  850] **********************************************************************************************       505
Bin[ 850 -  900] ***************************************************************************************************  536
Bin[ 900 -  950] *****************************************************************************************            478
Bin[ 950 - 1000] *******************************************************************************************          490
total count is 10000


```

### Installation

If you have a working go installation on a Unix-like OS:

> ```go get github.com/hotei/histogram```

Will copy github.com/hotei/histogram to the first entry of your $GOPATH

or if go is not installed yet :

> ```cd DestinationDirectory```

> ```git clone https://github.com/hotei/histogram.git```

### Features

* Very simple but flexible

### Limitations

* Only set up for integers at the moment but easy to adapt

### Usage

Typical usage is demonstrated in the test program.

_go test_ will show the results

### BUGS

? any ?

### To-Do

* Essential:
	* scalable X-11 version
* Nice:
	* vertical bars
	* Output to PNG/SVG
* Nice but no immediate need:
 * TBD

### Change Log

* 2015-06-15 Started
 * validated with 1.4.2
 
### Resources

* [go language reference] [1] 
* [go standard library package docs] [2]
* [Source for program] [3]

[1]: http://golang.org/ref/spec/ "go reference spec"
[2]: http://golang.org/pkg/ "go package docs"
[3]: http://github.com/hotei/program "github.com/hotei/histogram"

Comments can be sent to <hotei1352@gmail.com> or to user "hotei" at github.com.
License is BSD-two-clause, in file "LICENSE"

License
-------
The 'program' go package/program is distributed under the Simplified BSD License:

> Copyright (c) 2015 David Rook. All rights reserved.
> 
> Redistribution and use in source and binary forms, with or without modification, are
> permitted provided that the following conditions are met:
> 
>    1. Redistributions of source code must retain the above copyright notice, this list of
>       conditions and the following disclaimer.
> 
>    2. Redistributions in binary form must reproduce the above copyright notice, this list
>       of conditions and the following disclaimer in the documentation and/or other materials
>       provided with the distribution.
> 
> THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDER ``AS IS'' AND ANY EXPRESS OR IMPLIED
> WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
> FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> OR
> CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
> CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
> SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
> ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
> NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
> ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

Documentation (c) 2015 David Rook 

// EOF README.md  (this is a markdown document and tested OK with blackfriday)
