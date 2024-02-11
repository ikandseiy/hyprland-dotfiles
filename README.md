## Arch Linux Laptop Optimization Guide For Practical Use
https://www.reddit.com/r/archlinux/comments/rz6294/arch_linux_laptop_optimization_guide_for/


## App list
```bash
hyprland # tiling manager
hyprpaper # Wallpaper
vim (neovim) # cli editor
zsh # replace bash (also install oh-my-zsh)
intel-ucode/amd-ucode # microcode for intel/amd laptop
bluez # bluetooth
firefox # browser
dunst # Notification
sddm-git # Login manager
alacritty # Terminal
waybar # Top Bar
fastfetch # cli system info
brightnessctl # Screen brightness command line utils
pavucontrol # GUI pulseaudio controller
pamixer # Pulseaudio command line utils
pipewire # Audio
pipewire-pulse # bluetooth audio support
wireplumber # automatice choose profile
swaylock # lockscreen
grim # screenshot - Screenshot utility for Wayland
slurp # select region for grim
rofi-lbonn-wayland # application launcher
cliphist # clipboard
nautilus # File explorer
gvfs # Show Trash, Computer and other devices in thunar
gtk4 # Necessary for Chrome to use Fcitx5
sudo pacman -S $(pacman -Ssq noto-fonts) # fotns
gnome-keyring # Store secrets, passwords, keys, certificates
polkit-kde-agent # Authentication Agent
qt6-wayland # Hyprland Need
qt5-wayland # Hyprland Need
qt5ct # Hyprland Need
xdg-desktop-portal-gtk # Chrome needed, choose file & upload something
xdg-desktop-portal-hyprland-git # Screen Sharing
waybar-updates # Updates notification
tlp # power managment
tlpui # GUI for tlp
Thermald #Intel Only (sudo systemctl enable --now thermald)
nerd-fonts # Pack fonts - sudo pacman -S $(pacman -Ssq 'noto-fonts-*')

Optional:
  broadcom-wl-dkms # broadcom wireless driver (enable wifi - $sudo modprobe brcmfmac)
  visual-studio-code-bin # visual studio code
  telegram-desktop # telegram
  qbittorrent # torrent client

```
