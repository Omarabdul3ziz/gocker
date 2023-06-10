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

# What is a container?
- Process isolation?
- Processess sharing rescourse?
- bundling dependancies?
- ...

---
# Namespaces
What a process can see.
- pid
- mnt
- net
- uts
- ipc
- user
---
# Cgroups
What a process can use.
- memory
- cpu
- net
- ...
---
![img](./static/meme.png)