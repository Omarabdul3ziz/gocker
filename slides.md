---
marp: true
theme: dark-theme
---

<style>
:root {
    background: #1b1b1b;
    color: white;
    font-family: monospace;
}
h1 {
    color: white;
}
</style>

# Containers use-cases.
- Isolating Services
- Sharing Resources
- Bundling Dependencies
- ... and more.

---
# Namespaces
Isolate what a process can see.
- pid
- mnt
- net
- uts
- ipc
- user
---
# Cgroups
Limits what a process can use
- memory
- cpu
- pid
- i/o
- ...
---
![img](./static/meme.png)