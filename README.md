# gocker

A container in about 150 lines of Go, written for a talk on what a container actually is: a process with its own namespaces, limited by cgroups.

## How it works

`gocker run <cmd>` re-executes itself (`/proc/self/exe child <cmd>`) with new namespaces, and the child sets up the container before running the command:

1. **Namespaces**: the child starts in new UTS, PID and network namespaces (`CLONE_NEWUTS | CLONE_NEWPID | CLONE_NEWNET`).
2. **Cgroups**: it joins a `gocker` cgroup capped at 10 MB of memory and 20 processes.
3. **Hostname**: it sets the hostname to `gocker`, visible only inside the UTS namespace.
4. **Root filesystem**: it `chroot`s into a prepared root filesystem and mounts a fresh `/proc`, so `ps` only shows the container's own processes.
5. **Run**: it runs the command, then unmounts `/proc`.

## Run it

Needs Linux, root, and cgroup v1 (`/sys/fs/cgroup/memory` and `/sys/fs/cgroup/pids`).

```bash
# a root filesystem to chroot into, e.g. from an Alpine minirootfs
sudo mkdir -p /home/omar/images/layer
sudo tar -xzf alpine-minirootfs-*.tar.gz -C /home/omar/images/layer

go build -o gocker .
sudo ./gocker run /bin/sh
```

The rootfs path is hard-coded in `child()` in `main.go`. Change it to wherever yours lives.

## The talk

[`slides.md`](slides.md) is the deck (Marp): container use cases, namespaces, cgroups, and the demo above. [`note.md`](note.md) has the notes behind it.
