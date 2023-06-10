# Namespaces
- a way to limit what a process can see. illusion its the only process running on the host.
- isolate the system resouces between independant processes. twor processes run on the same host have same ID if both in different `PID` namespace.
- it is useful to make each process care its own business without affecting other processes. So process A in a different name space of B. it is not aware of it's existence.
- `unshare` runs program with namespaces unshared from the parent.
Example on `UTS` namespace. isolates the hostname.
```bash
unshare --help

hostname
/bin/bash
hostname 

# still the same. inhereted form the parent
unshare --uts /bin/bash
hostname

ip link

unshare --net /bin/bash
ip link


# check the namespace 
readlink /proc/self/ns/uts
```