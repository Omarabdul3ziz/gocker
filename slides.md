---
marp: true
theme: dark-theme
style: |
    img[alt~="center"] {
      display: block;
      margin: 0 auto;
    }
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
This is basically what a container is a process with new set of namespaces and limitation with cgroups.

<br>


![w:500 center](./static/meme.png)