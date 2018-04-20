# sha3sum

A SHA3 hash calculator

# Usage

sha3sum file  
sha3sum [-l length] [-q] file  
sha3sum -h  
sha3sum --help  

l = "length", value can be 224, 256, 384 or 512  
q = "quiet", output the hash only  
file = the file to be hashed  

## Examples

```sh
$ sha3sum myfile
$ sha3sum -l 384 myfile
$ sha3sum -l 384 -q myfile
$ sha3sum -h
```
That's pretty much all there is.



License
----

Do whatever you like with it. Sell it and get rich. Send me some money if you're lucky.

