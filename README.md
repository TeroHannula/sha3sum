sha3sum - a SHA3 hash calulator for 224, 256, 384 and 512 bit hashes.

Usage: 

sha3sum file
sha3sum [-l length] [-q] file
sha3sum -h 
sha3sum --help

l = "length", value can be 224, 256, 384 or 512
q = "quiet", output the hash only
file = the file to be hashed

For example:

sha3sum -l 384 -q myfile

That's pretty much all there is.
