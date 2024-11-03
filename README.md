# kanban-dunkan

## Linux Dependencies

### Ubuntu/Debian
```bash
sudo apt-get update
sudo apt-get install libx11-dev libxrandr-dev libxinerama-dev \
libxcursor-dev libxi-dev libgl1-mesa-dev \
libxxf86vm-dev
```

### Fedora/RHEL
```bash
sudo dnf install libX11-devel libXrandr-devel libXinerama-devel \
libXcursor-devel libXi-devel mesa-libGL-devel \
libXxf86vm-devel
```

### Arch Linux
```bash
sudo pacman -S libx11 libxrandr libxinerama libxcursor libxi \
mesa libxxf86vm
```